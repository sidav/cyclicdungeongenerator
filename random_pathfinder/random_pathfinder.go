package random_pathfinder

import "github.com/sidav/golibrl/random"

const (
	DIAGONAL_COST = 14
	STRAIGHT_COST = 10
	HEURISTIC_MULTIPLIER = 10
	MAX_PATHFINDING_STEPS = 175 // Increase in case of stupid pathfinding. Decrease in case of lag.
)

type Cell struct {
	X, Y            int
	g            int
	costToMoveThere int
	parent          *Cell
	Child           *Cell
}

func (c *Cell) GetCoords() (int, int) {
	return c.X, c.Y
}

func (c *Cell) setG(inc int) {
	if c.parent != nil {
		c.g = c.parent.g + inc
	}
}

func (c *Cell) GetNextStepVector() (int, int) {
	var x, y int
	if c.Child != nil {
		x = c.Child.X - c.X
		y = c.Child.Y - c.Y
	}
	return x, y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getIndexOfRandomCellFromList(List *[]*Cell) int {
	return random.Random(len(*List))
}

func (c *Cell) setChildsForPath() {
	// path := make([]*Cell, 0)
	curcell := c
	for curcell.parent != nil {
		// path = append(path, curcell)
		curcell.parent.Child = curcell
		curcell = curcell.parent
	}
	return
}

func FindPath(costMap *[][]int, fromx, fromy, tox, toy int, diagonalMoveAllowed, forceGetPath, forceIncludeFinish bool) *Cell {
	openList := make([]*Cell, 0)
	closedList := make([]*Cell, 0)
	var currentCell *Cell
	total_steps := 0
	targetReached := false

	// step 1
	origin := &Cell{X: fromx, Y: fromy, costToMoveThere: 0}
	openList = append(openList, origin)
	// step 2
	for !targetReached {
		// sub-step 2a:
		currentCellIndex := getIndexOfRandomCellFromList(&openList)
		currentCell = openList[currentCellIndex]
		// sub-step 2b:
		closedList = append(closedList, currentCell)
		openList = append(openList[:currentCellIndex], openList[currentCellIndex+1:]...) // this friggin' magic removes currentCellIndex'th element from openList
		//sub-step 2c:
		analyzeNeighbors(currentCell, &openList, &closedList, costMap, tox, toy, diagonalMoveAllowed, forceIncludeFinish)
		//sub-step 2d:
		total_steps += 1
		targetInOpenList := getCellWithCoordsFromList(&openList, tox, toy)
		if targetInOpenList != nil {
			currentCell = targetInOpenList
			currentCell.setChildsForPath()
			return origin
		}
		if len(openList) == 0 || total_steps > MAX_PATHFINDING_STEPS {
			if forceGetPath { // makes the routine always return path to the closest possible cell to (tox, toy) even if the precise path does not exist.
				currentCell = closedList[getIndexOfRandomCellFromList(&closedList)]
				currentCell.setChildsForPath()
				return origin
			} else {
				return nil
			}
		}
	}
	return nil
}

func analyzeNeighbors(curCell *Cell, openlist *[]*Cell, closedlist *[]*Cell, costMap *[][]int, targetX, targetY int, diagAllowed, forceIncludeFinish bool) {
	cost := 0
	cx, cy := curCell.X, curCell.Y
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if (i == 0 && j == 0) || (!diagAllowed && i != 0 && j != 0) {
				continue
			}
			x, y := cx+i, cy+j
			if areCoordsValidForCostMap(x, y, costMap) {
				// if (x != targetX || y != targetY) &&
				if (*costMap)[x][y] == -1 || getCellWithCoordsFromList(closedlist, x, y) != nil { // Cell is impassable or is in closed list
					if !(forceIncludeFinish && x == targetX && y == targetY) { // if forceIncludeFinish is true, then we won't ignore finish cell whether it is passable or whatever.
						continue // ignore it
					}
				}
				// TODO: add actual "cost to move there" from costMap
				if (i * j) != 0 { // the Cell under consideration is lying diagonally
					cost = DIAGONAL_COST
				} else {
					cost = STRAIGHT_COST
				}
				curNeighbor := getCellWithCoordsFromList(openlist, x, y)
				if curNeighbor != nil {
					if curNeighbor.g > curCell.g+cost {
						curNeighbor.parent = curCell
						curNeighbor.setG(cost)
					}
				} else {
					curNeighbor = &Cell{X: x, Y: y, parent: curCell}
					curNeighbor.setG(cost)
					*openlist = append(*openlist, curNeighbor)
				}
			}
		}
	}
}

func getCellWithCoordsFromList(list *[]*Cell, x, y int) *Cell {
	for _, c := range *list {
		if c.X == x && c.Y == y {
			return c
		}
	}
	return nil
}

func areCoordsValidForCostMap(x, y int, costMap *[][]int) bool {
	return x >= 0 && y >= 0 && (x < len(*costMap)) && (y < len((*costMap)[0]))
}
