package layout_generation

import "fmt"

func (p *pattern) ShowInitialAndOptimizedInstructionOrders() {
	fmt.Printf("=== OPTIMIZATION OF PATTERN ===\n")
	fmt.Printf("=== BEFORE ===\n")
	for i := range p.instructions {
		fmt.Printf("%d: %s\n", i, p.instructions[i].instructionText)
		p.instructions[i].instructionText += fmt.Sprintf(" (old num %d)", i)
	}
	p.optimizeStepsOrder()
	fmt.Printf("\n=== AFTER ===\n")
	for i := range p.instructions {
		fmt.Printf("%d: %s\n", i, p.instructions[i].instructionText)
	}
	fmt.Printf("===+===+===\n")
}

func (p *pattern) optimizeStepsOrder() {
	// move all "add node" to the beginning
	instructionMoved := true
	for instructionMoved {
		instructionMoved = false
		instructionNeedsToMove := false
	iterateInstructionsUp:
		for i := range p.instructions {
			types := []int{
				ACTION_PLACE_NODE_AT_EMPTY,
				ACTION_PLACE_NODE_AT_PATH,
				ACTION_PLACE_NODE_NEAR_PATH,
				ACTION_PLACE_OBSTACLE_AT_COORDS,
			}
			checkedType := p.instructions[i].actionType
			for _, typeFromList := range types {
				if checkedType == typeFromList {
					if instructionNeedsToMove {
						p.moveInstructionToBeginningOrToCode(i, types)
						instructionMoved = true
						break iterateInstructionsUp
					}
				}
			}
		}
	}
	// move all non-creating instructions to end
	instructionMoved = true
	for instructionMoved {
		instructionMoved = false
	iterateInstructionsDown:
		for i := range p.instructions {
			types := []int{
				ACTION_PLACE_NODE_AT_EMPTY,
				ACTION_PLACE_NODE_AT_PATH,
				ACTION_PLACE_NODE_NEAR_PATH,
				ACTION_PLACE_OBSTACLE_AT_COORDS,
			}
			checkedType := p.instructions[i].actionType
			for _, typeFromList := range types {
				if checkedType == typeFromList {
					break iterateInstructionsDown
				}
			}
			p.moveInstructionToEnd(i)
			instructionMoved = true
			break
		}
	}
	// move each "add path" as up as possible
	for i := range p.instructions {
		for j := i + 1; j < len(p.instructions); j++ {
			areNodesPlaced := false
			if p.instructions[j].actionType == ACTION_PLACE_PATH_FROM_TO {
				areNodesPlaced = p.areNodesPlacedUntilStep(p.instructions[j].nameFrom, p.instructions[j].nameTo, i)
				if areNodesPlaced {
					p.moveInstructionUpFromTo(j, i)
				}
			}
		}
	}
	// move each lock in order of lockId
	swapped := true
	for swapped {
		swapped = false
		for i := range p.instructions {
			if p.instructions[i].actionType != ACTION_SET_NODE_CONNECTION_LOCKED_FROM_PATH {
				continue
			}
			for j := i + 1; j < len(p.instructions); j++ {
				if p.instructions[j].actionType == ACTION_SET_NODE_CONNECTION_LOCKED_FROM_PATH {
					if p.instructions[i].lockNumber > p.instructions[j].lockNumber {
						p.swapInstructionsAtIndices(i, j)
						swapped = true
					}
				}
			}
		}
	}
}

func (p *pattern) areNodesPlacedUntilStep(name1, name2 string, step int) bool {
	firstPlaced := false
	secondPlaced := false
	for i := 0; i < step; i++ {
		if p.instructions[i].actionType == ACTION_PLACE_NODE_AT_EMPTY ||
			p.instructions[i].actionType == ACTION_PLACE_NODE_AT_PATH ||
			p.instructions[i].actionType == ACTION_PLACE_NODE_NEAR_PATH {
			if p.instructions[i].nameOfNode == name1 {
				firstPlaced = true
			}
			if p.instructions[i].nameOfNode == name2 {
				secondPlaced = true
			}
			if firstPlaced && secondPlaced {
				return true
			}
		}
	}
	return false
}

func (p *pattern) swapInstructionsAtIndices(i, j int) {
	t := p.instructions[j]
	p.instructions[j] = p.instructions[i]
	p.instructions[i] = t
}

func (p *pattern) moveInstructionUpFromTo(from, to int) {
	if to > from {
		panic("to should be < from!")
	}
	for x := from; x > to; x-- {
		p.swapInstructionsAtIndices(x, x-1)
	}
}

func (p *pattern) moveInstructionToBeginningOrToCode(from int, codes []int) {
	for x := from; x > 0; x-- {
		checkedType := p.instructions[x-1].actionType
		for _, typeFromList := range codes {
			if checkedType == typeFromList {
				break
			}
		}
		p.swapInstructionsAtIndices(x, x-1)
	}
}

func (p *pattern) moveInstructionToEnd(from int) {
	for x := from; x < len(p.instructions)-1; x++ {
		p.swapInstructionsAtIndices(x, x+1)
	}
}
