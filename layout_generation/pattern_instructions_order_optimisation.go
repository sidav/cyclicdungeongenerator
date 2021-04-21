package layout_generation

func (p *pattern) optimizeStepsOrder() {
	//fmt.Println("BEFORE")
	//for i := range p.instructions {
	//	fmt.Println(p.instructions[i].instructionText)
	//}
	// first, move all "add node" to the beginning
	for i := range p.instructions {
		for j := i + 1; j < len(p.instructions); j++ {
			if p.instructions[i].actionType == ACTION_PLACE_NODE_AT_EMPTY ||
				p.instructions[i].actionType == ACTION_PLACE_NODE_AT_PATH ||
				p.instructions[i].actionType == ACTION_PLACE_NODE_NEAR_PATH ||
				p.instructions[i].actionType == ACTION_PLACE_OBSTACLE_AT_COORDS {
				break 
			}
			p.swapInstructionsAtIndices(i, j)
		}
	}
	// second, move each "add path" as up as possible
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
	//fmt.Println("AFTER")
	//for i := range p.instructions {
	//	fmt.Println(p.instructions[i].instructionText)
	//}
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
