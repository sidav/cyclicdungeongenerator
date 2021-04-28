package layout_generation

import (
	"CyclicDungeonGenerator/random"
)

type LayoutMap struct {
	// room map with connections or smth
	elements       [][]*element
	rnd            *random.FibRandom
	randomizePaths bool
}

func (lm *LayoutMap) init(sizex, sizey int, rnd *random.FibRandom, randomizePaths bool) {
	lm.elements = make([][]*element, sizex)
	for i := range lm.elements {
		lm.elements[i] = make([]*element, sizey)
	}
	for x := 0; x < sizex; x++ {
		for y := 0; y < sizey; y++ {
			lm.elements[x][y] = &element{connections: map[string]*connection{"north": nil, "east": nil, "south": nil, "west": nil}}
		}
	}
	lm.rnd = rnd
	lm.randomizePaths = randomizePaths
}

func (lm *LayoutMap) randomizeTagLocationsPerNode() {
	w, h := lm.GetSize()
	var alreadyCheckedNodes []string
	// randomize tags for multi-cell nodes
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			if lm.elements[x][y].IsNode() {
				name := lm.elements[x][y].GetName()
				// avoid unnecessary loops
				if name == "" || lm.countNodesOfName(name) < 2 {
					continue
				}
				// check if the name was already randomized
				for _, chckname := range alreadyCheckedNodes {
					if name == chckname {
						continue
					}
				}
				tag := ""
				nodesWithName := lm.getAllCoordsOfNode(name)
				// clear the tag from everything
				for _, n := range nodesWithName {
					nde := lm.elements[n[0]][n[1]]
					if tag == "" {
						tag = nde.nodeInfo.nodeTag
						nde.nodeInfo.nodeTag = ""
					}
				}
				// assign the tag to random node of list
				randNodeIndex := lm.rnd.Rand(len(nodesWithName))
				i, j := nodesWithName[randNodeIndex][0], nodesWithName[randNodeIndex][1]
				lm.elements[i][j].nodeInfo.nodeTag = tag
				alreadyCheckedNodes = append(alreadyCheckedNodes, name)
			}
		}
	}
}

func (lm *LayoutMap) placeNodeAtCoords(x, y int, nodeName string) {
	lm.elements[x][y].nodeInfo = &nodeCell{nodeName: nodeName}
}

func (lm *LayoutMap) placePathAtCoords(x, y int, pathNum int) {
	lm.elements[x][y].pathInfo = &pathCell{pathNum}
}

func (lm *LayoutMap) placeObstacleAtCoords(x, y int) {
	lm.elements[x][y].isObstacle = true
}

func (lm *LayoutMap) countEmptyElements() int {
	w, h := lm.GetSize()
	empties := 0
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			if lm.elements[x][y].isEmpty() {
				empties++
			}
		}
	}
	return empties
}

func (lm *LayoutMap) countNodesOfName(nodeName string) int {
	w, h := lm.GetSize()
	count := 0
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			if lm.elements[x][y].IsNode() && lm.elements[x][y].GetName() == nodeName {
				count++
			}
		}
	}
	return count
}

func (lm *LayoutMap) removeAllObstacles() {
	for x := 0; x < len(lm.elements); x++ {
		for y := 0; y < len(lm.elements[0]); y++ {
			lm.elements[x][y].isObstacle = false
		}
	}
}

func (lm *LayoutMap) getRandomPathCoordsAndRandomCellNearPath(pathNum int, allowNearNode bool) (int, int, int, int) {
	const tries = 10
	for try := 0; try < tries; try++ {
		px, py := lm.getRandomPathCellCoords(pathNum, allowNearNode)
		if px == -1 && py == -1 {
			continue
		}
		for try2 := 0; try2 < tries; try2++ {
			x, y := lm.rnd.RandInRange(px-1, px+1), lm.rnd.RandInRange(py-1, py+1)
			if (px-x)*(py-y) != 0 { // diagonal direction is restricted
				continue
			}
			if lm.areCoordsValid(x, y) && lm.elements[x][y].isEmpty() {
				return px, py, x, y
			}
		}
	}
	return -1, -1, -1, -1
}

func (lm *LayoutMap) getRandomNonEmptyCoordsAndRandomCellNearIt() (int, int, int, int) {
	px, py := lm.getRandomNonEmptyCellCoords(1)
	if px == -1 && py == -1 {
		return -1, -1, -1, -1
	}
	x, y := lm.getRandomEmptyCellNearCoords(px, py)
	if x == -1 && y == -1 {
		return -1, -1, -1, -1
	}
	return px, py, x, y
}

//func (lm *LayoutMap) getRandomEmptyCellCoords(minEmptyCellsNear int, cornerAllowed, edgeAllowed bool) (int, int) {
//	w, h := lm.GetSize()
//	emptiesX := make([]int, 0)
//	emptiesY := make([]int, 0)
//	for x := 0; x < w; x++ {
//		for y := 0; y < h; y++ {
//			corner := (x == 0 && y == 0) || (x == 0 && y == h-1) || (x == w-1 && y == 0) || (x == w-1 && y == h-1)
//			edge := x*y == 0 || x == w-1 || y == h-1
//			if (!cornerAllowed && corner) || (!edgeAllowed && edge) {
//				continue
//			}
//			if lm.elements[x][y].isEmpty() && (lm.countEmptyCellsNear(x, y) >= minEmptyCellsNear) {
//				emptiesX = append(emptiesX, x)
//				emptiesY = append(emptiesY, y)
//			}
//		}
//	}
//	if len(emptiesX) == 0 {
//		return -1, -1
//	}
//	index := lm.rnd.Rand(len(emptiesX))
//	return emptiesX[index], emptiesY[index]
//}

func (lm *LayoutMap) getRandomEmptyCellCoords(minEmptyCellsNear int, cornerAllowed, edgeAllowed bool, NodeDistMap *map[string]int) (int, int) {
	w, h := lm.GetSize()
	emptiesX := make([]int, 0)
	emptiesY := make([]int, 0)
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			corner := (x == 0 && y == 0) || (x == 0 && y == h-1) || (x == w-1 && y == 0) || (x == w-1 && y == h-1)
			edge := x*y == 0 || x == w-1 || y == h-1
			if (!cornerAllowed && corner) || (!edgeAllowed && edge) {
				continue
			}

			if lm.elements[x][y].isEmpty() && (lm.countEmptyCellsNear(x, y) >= minEmptyCellsNear) {
				coordsAreInDist := true
				for nodeName, minDist := range *NodeDistMap {
					currNodeCoordsForDistCheck := lm.getAllCoordsOfNode(nodeName)
					if len(currNodeCoordsForDistCheck) == 0 {
						continue
					}
					atLeastOneIsGood := false
					for i := range currNodeCoordsForDistCheck {
						if euclideanDistance(currNodeCoordsForDistCheck[i][0], currNodeCoordsForDistCheck[i][1], x, y) >= minDist {
							atLeastOneIsGood = true
						}
					}
					if !atLeastOneIsGood {
						coordsAreInDist = false
						break
					}
				}
				if coordsAreInDist {
					emptiesX = append(emptiesX, x)
					emptiesY = append(emptiesY, y)
				}
			}
		}
	}
	if len(emptiesX) == 0 {
		return -1, -1
	}
	index := lm.rnd.Rand(len(emptiesX))
	return emptiesX[index], emptiesY[index]
}

// creates additional node with the same name.
func (lm *LayoutMap) tryGrowingNodeByName(nodeName string) {
	x, y := lm.getAnyOfCoordsOfNode(nodeName)
	lm.tryGrowingNodeFromCoords(x, y)
}

// creates additional node with the same name.
func (lm *LayoutMap) growAllNodesToFillSpace(maxNodeSize int) {
	w, h := lm.GetSize()
	currentEmpty := lm.countEmptyElements()
	prevEmpty := -1
	for currentEmpty > 0 && prevEmpty != currentEmpty {
		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				if lm.elements[x][y].IsNode() {
					nodeName := lm.elements[x][y].GetName()
					size := lm.countNodesOfName(nodeName)
					if size < maxNodeSize || maxNodeSize == 0 {
						lm.tryGrowingNodeFromCoords(x, y)
					}
				}
			}
		}
		prevEmpty = currentEmpty
		currentEmpty = lm.countEmptyElements()
	}
}

func (lm *LayoutMap) tryGrowingNodeFromCoords(x, y int) {
	if lm.elements[x][y].IsNode() {
		ex, ey := lm.getRandomEmptyCellNearCoords(x, y)
		if ex > -1 && ey > -1 {
			lm.placeNodeAtCoords(ex, ey, lm.elements[x][y].GetName())
			lm.setConnectionsBetweenTwoCoords(&connection{IsNodeExtension: true}, x, y, ex, ey)
		}
	}
}

func (lm *LayoutMap) setConnectionsBetweenTwoCoords(c *connection, x1, y1, x2, y2 int) {
	// TODO: check for adjacency
	lm.elements[x1][y1].setConnectionByCoords(c, x2-x1, y2-y1)
	lm.elements[x2][y2].setConnectionByCoords(c, x1-x2, y1-y2)
}

func (lm *LayoutMap) getRandomEmptyCellCoordsInRange(fx, fy, tx, ty, minEmptyCellsNear int) (int, int) { // range inclusive
	emptiesX := make([]int, 0)
	emptiesY := make([]int, 0)
	w, h := lm.GetSize()
	if fx > tx {
		t := tx
		tx = fx
		fx = t
	}
	if fy > ty {
		t := ty
		ty = fy
		fy = t
	}
	if tx >= w {
		tx = w - 1
	}
	if ty >= h {
		ty = h - 1
	}
	for x := fx; x <= tx; x++ {
		for y := fy; y <= ty; y++ {
			if lm.elements[x][y].isEmpty() && (lm.countEmptyCellsNear(x, y) >= minEmptyCellsNear) {
				emptiesX = append(emptiesX, x)
				emptiesY = append(emptiesY, y)
			}
		}
	}
	if len(emptiesX) == 0 {
		return -1, -1
	}
	index := lm.rnd.Rand(len(emptiesX))
	return emptiesX[index], emptiesY[index]
}

func (lm *LayoutMap) getRandomEmptyCellNearCoords(nx, ny int) (int, int) {
	emptiesX := make([]int, 0)
	emptiesY := make([]int, 0)
	for x := nx - 1; x <= nx+1; x++ {
		for y := ny - 1; y <= ny+1; y++ {
			if (nx-x)*(ny-y) != 0 { // restrict diagonals
				continue
			}
			if (x != nx || y != ny) && lm.areCoordsValid(x, y) && lm.elements[x][y].isEmpty() {
				emptiesX = append(emptiesX, x)
				emptiesY = append(emptiesY, y)
			}
		}
	}
	if len(emptiesX) == 0 {
		return -1, -1
	}
	index := lm.rnd.Rand(len(emptiesX))
	return emptiesX[index], emptiesY[index]
}

func (lm *LayoutMap) getRandomNonEmptyCellCoords(minEmptyCellsNear int) (int, int) {
	nonEmptiesX := make([]int, 0)
	nonEmptiesY := make([]int, 0)
	for x := 0; x < len(lm.elements); x++ {
		for y := 0; y < len(lm.elements[0]); y++ {
			// obstacles should not be counted as "non-empty"
			if !lm.elements[x][y].isObstacle && !lm.elements[x][y].isEmpty() && lm.countEmptyCellsNear(x, y) >= minEmptyCellsNear {
				nonEmptiesX = append(nonEmptiesX, x)
				nonEmptiesY = append(nonEmptiesY, y)
			}
		}
	}
	if len(nonEmptiesX) == 0 {
		return -1, -1
	}
	index := lm.rnd.Rand(len(nonEmptiesX))
	return nonEmptiesX[index], nonEmptiesY[index]
}

func (lm *LayoutMap) getRandomPathCellCoords(desiredPathNum int, nodesAllowed bool) (int, int) { // desiredPathNum -1 means any path
	pathsX := make([]int, 0)
	pathsY := make([]int, 0)
	for x := 0; x < len(lm.elements); x++ {
		for y := 0; y < len(lm.elements[0]); y++ {
			if !lm.elements[x][y].isPartOfAPath() {
				continue
			}
			if !nodesAllowed && lm.elements[x][y].nodeInfo != nil { // don't take nodes unless allowed
				continue
			}
			if desiredPathNum > -1 && desiredPathNum != lm.elements[x][y].pathInfo.pathNumber { // don't take cells of non-desired path numbers
				continue
			}
			pathsX = append(pathsX, x)
			pathsY = append(pathsY, y)
		}
	}
	if len(pathsX) == 0 {
		return -1, -1
	}
	index := lm.rnd.Rand(len(pathsX))
	return pathsX[index], pathsY[index]
}

func (lm *LayoutMap) areCoordsEmpty(x, y int) bool {
	return lm.elements[x][y].isEmpty()
}

func (lm *LayoutMap) isPathPresentAtCoords(x, y int) bool {
	return lm.elements[x][y].isPartOfAPath()
}

func (lm *LayoutMap) areCoordsEmptyOrPathOnly(x, y int) bool {
	return lm.elements[x][y].isPathOrEmpty()
}

func (lm *LayoutMap) countEmptyCellsNear(x, y int) int {
	count := 0
	w, h := lm.GetSize()
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i*j != 0 || i == 0 && j == 0 {
				continue
			}
			if x+i < 0 || x+i >= w || y+j < 0 || y+j >= h {
				continue
			}
			if lm.elements[x+i][y+j].isEmpty() {
				count++
			}
		}
	}
	return count
}

func (lm *LayoutMap) getAnyOfCoordsOfNode(nodeName string) (int, int) {
	possibleCoords := make([][2]int, 0)
	for x := 0; x < len(lm.elements); x++ {
		for y := 0; y < len(lm.elements[0]); y++ {
			if lm.elements[x][y].IsNode() && lm.elements[x][y].nodeInfo.nodeName == nodeName {
				possibleCoords = append(possibleCoords, [2]int{x, y})
			}
		}
	}
	if len(possibleCoords) > 0 {
		rndIndex := lm.rnd.Rand(len(possibleCoords))
		return possibleCoords[rndIndex][0], possibleCoords[rndIndex][1]
	}
	panic("getAnyOfCoordsOfNode failed with node " + nodeName)
	return -1, -1
}

func (lm *LayoutMap) getAllCoordsOfNode(nodeName string) [][2]int {
	possibleCoords := make([][2]int, 0)
	for x := 0; x < len(lm.elements); x++ {
		for y := 0; y < len(lm.elements[0]); y++ {
			if lm.elements[x][y].IsNode() && lm.elements[x][y].nodeInfo.nodeName == nodeName {
				possibleCoords = append(possibleCoords, [2]int{x, y})
			}
		}
	}
	if len(possibleCoords) > 0 {
		return possibleCoords
	}
	return [][2]int{}
	panic("getAllCoordsOfNode failed with node " + nodeName)
}

func (lm *LayoutMap) setAllNodeConnectionsLockedForPath(nodex, nodey, pathNum int, lockNum int) {
	e := lm.GetElement(nodex, nodey)
	for dir, v := range e.connections {
		if v != nil && v.pathNum == pathNum {
			v.IsLocked = true
			if v.LockNum < lockNum {
				v.LockNum = lockNum
			}
			// set the lock for neighbouring connection
			var dirCoords [2]int
			var opposingPath string
			if dir == "north" {
				dirCoords = [2]int{0, -1}
				opposingPath = "south"
			}
			if dir == "south" {
				dirCoords = [2]int{0, 1}
				opposingPath = "north"
			}
			if dir == "east" {
				dirCoords = [2]int{1, 0}
				opposingPath = "west"
			}
			if dir == "west" {
				dirCoords = [2]int{-1, 0}
				opposingPath = "east"
			}
			elem2 := lm.GetElement(nodex+dirCoords[0], nodey+dirCoords[1])
			elem2.connections[opposingPath].IsLocked = true
			if elem2.connections[opposingPath].LockNum < lockNum {
				elem2.connections[opposingPath].LockNum = lockNum
			}
		}
	}
}

func (lm *LayoutMap) areCoordsValid(x, y int) bool {
	w, h := lm.GetSize()
	return x >= 0 && x < w && y >= 0 && y < h
}

// exported

func (lm *LayoutMap) GetSize() (int, int) {
	return len(lm.elements), len(lm.elements[0])
}

func (lm *LayoutMap) GetElement(x, y int) *element {
	return lm.elements[x][y]
}

func (lm *LayoutMap) getPassabilityMapForPathfinder(pathsArePassable bool) *[][]int {
	const (
		minRandomCostIncrease = -100
		maxRandomCostIncrease = 10000
	)
	layoutWidth, layoutHeight := lm.GetSize()
	pmap := make([][]int, layoutWidth)
	for i := range pmap {
		pmap[i] = make([]int, layoutHeight)
	}

	for x := 0; x < layoutWidth; x++ {
		for y := 0; y < layoutHeight; y++ {
			if lm.areCoordsEmpty(x, y) || pathsArePassable && lm.areCoordsEmptyOrPathOnly(x, y) {
				pmap[x][y] = 1
				if lm.isPathPresentAtCoords(x, y) {
					pmap[x][y] += maxRandomCostIncrease
				}
				// TODO: think how to better randomize path costs
				if lm.randomizePaths {
					// lowering the "from" increases path randomness, but also makes the generator to fail more frequently
					// because it increases the probability for creating a non-existing path
					// "* 10" is to compensate the heuristics in the pathfinder
					pmap[x][y] += lm.rnd.RandInRange(minRandomCostIncrease, maxRandomCostIncrease) * 10
				}
			} else {
				pmap[x][y] = -1
			}
		}
	}
	return &pmap
}
