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

func (r *LayoutMap) init(sizex, sizey int, rnd *random.FibRandom, randomizePaths bool) {
	r.elements = make([][]*element, sizex)
	for i := range r.elements {
		r.elements[i] = make([]*element, sizey)
	}
	for x := 0; x < sizex; x++ {
		for y := 0; y < sizey; y++ {
			r.elements[x][y] = &element{connections: map[string]*connection{"north": nil, "east": nil, "south": nil, "west": nil}}
		}
	}
	r.rnd = rnd
	r.randomizePaths = randomizePaths
}

func (r *LayoutMap) placeNodeAtCoords(x, y int, nodeName string) {
	r.elements[x][y].nodeInfo = &nodeCell{nodeName: nodeName}
}

func (r *LayoutMap) placePathAtCoords(x, y int, pathNum int) {
	r.elements[x][y].pathInfo = &pathCell{pathNum}
}

func (r *LayoutMap) placeObstacleAtCoords(x, y int) {
	r.elements[x][y].isObstacle = true
}

func (r *LayoutMap) countEmptyElements() int {
	w, h := r.GetSize()
	empties := 0
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			if r.elements[x][y].isEmpty() {
				empties++
			}
		}
	}
	return empties
}

func (r *LayoutMap) countNodesOfName(nodeName string) int {
	w, h := r.GetSize()
	count := 0
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			if r.elements[x][y].IsNode() && r.elements[x][y].GetName() == nodeName {
				count++
			}
		}
	}
	return count
}

func (r *LayoutMap) removeAllObstacles() {
	for x := 0; x < len(r.elements); x++ {
		for y := 0; y < len(r.elements[0]); y++ {
			r.elements[x][y].isObstacle = false
		}
	}
}

func (r *LayoutMap) getRandomPathCoordsAndRandomCellNearPath(pathNum int, allowNearNode bool) (int, int, int, int) {
	const tries = 10
	for try := 0; try < tries; try++ {
		px, py := r.getRandomPathCellCoords(pathNum, allowNearNode)
		if px == -1 && py == -1 {
			continue
		}
		for try2 := 0; try2 < tries; try2++ {
			x, y := r.rnd.RandInRange(px-1, px+1), r.rnd.RandInRange(py-1, py+1)
			if (px-x)*(py-y) != 0 { // diagonal direction is restricted
				continue
			}
			if r.areCoordsValid(x, y) && r.elements[x][y].isEmpty() {
				return px, py, x, y
			}
		}
	}
	return -1, -1, -1, -1
}

func (r *LayoutMap) getRandomNonEmptyCoordsAndRandomCellNearIt() (int, int, int, int) {
	px, py := r.getRandomNonEmptyCellCoords(1)
	if px == -1 && py == -1 {
		return -1, -1, -1, -1
	}
	x, y := r.getRandomEmptyCellNearCoords(px, py)
	if x == -1 && y == -1 {
		return -1, -1, -1, -1
	}
	return px, py, x, y
}

func (r *LayoutMap) getRandomEmptyCellCoords(minEmptyCellsNear int, cornerAllowed, edgeAllowed bool) (int, int) {
	w, h := r.GetSize()
	emptiesX := make([]int, 0)
	emptiesY := make([]int, 0)
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			corner := (x == 0 && y == 0) || (x == 0 && y == h-1) || (x == w-1 && y == 0) || (x == w-1 && y == h-1)
			edge := x*y == 0 || x == w-1 || y == h-1
			if (!cornerAllowed && corner) || (!edgeAllowed && edge) {
				continue
			}
			if r.elements[x][y].isEmpty() && (r.countEmptyCellsNear(x, y) >= minEmptyCellsNear) {
				emptiesX = append(emptiesX, x)
				emptiesY = append(emptiesY, y)
			}
		}
	}
	if len(emptiesX) == 0 {
		return -1, -1
	}
	index := r.rnd.Rand(len(emptiesX))
	return emptiesX[index], emptiesY[index]
}

// creates additional node with the same name.
func (r *LayoutMap) tryGrowingNodeByName(nodeName string) {
	x, y := r.getAnyOfCoordsOfNode(nodeName)
	r.tryGrowingNodeFromCoords(x, y)
}

// creates additional node with the same name.
func (r *LayoutMap) growAllNodesToFillSpace(maxNodeSize int) {
	w, h := r.GetSize()
	currentEmpty := r.countEmptyElements()
	prevEmpty := -1
	for currentEmpty > 0 && prevEmpty != currentEmpty {
		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				if r.elements[x][y].IsNode() {
					nodeName := r.elements[x][y].GetName()
					size := r.countNodesOfName(nodeName)
					if size < maxNodeSize || maxNodeSize == 0 {
						r.tryGrowingNodeFromCoords(x, y)
					}
				}
			}
		}
		prevEmpty = currentEmpty
		currentEmpty = r.countEmptyElements()
	}
}

func (r *LayoutMap) tryGrowingNodeFromCoords(x, y int) {
	if r.elements[x][y].IsNode() {
		ex, ey := r.getRandomEmptyCellNearCoords(x, y)
		if ex > -1 && ey > -1 {
			r.placeNodeAtCoords(ex, ey, r.elements[x][y].GetName())
			r.setConnectionsBetweenTwoCoords(&connection{IsNodeExtension: true}, x, y, ex, ey)
		}
	}
}

func (r *LayoutMap) setConnectionsBetweenTwoCoords(c *connection, x1, y1, x2, y2 int) {
	// TODO: check for adjacency
	r.elements[x1][y1].setConnectionByCoords(c, x2-x1, y2-y1)
	r.elements[x2][y2].setConnectionByCoords(c, x1-x2, y1-y2)
}

func (r *LayoutMap) getRandomEmptyCellCoordsInRange(fx, fy, tx, ty, minEmptyCellsNear int) (int, int) { // range inclusive
	emptiesX := make([]int, 0)
	emptiesY := make([]int, 0)
	w, h := r.GetSize()
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
			if r.elements[x][y].isEmpty() && (r.countEmptyCellsNear(x, y) >= minEmptyCellsNear) {
				emptiesX = append(emptiesX, x)
				emptiesY = append(emptiesY, y)
			}
		}
	}
	if len(emptiesX) == 0 {
		return -1, -1
	}
	index := r.rnd.Rand(len(emptiesX))
	return emptiesX[index], emptiesY[index]
}

func (r *LayoutMap) getRandomEmptyCellNearCoords(nx, ny int) (int, int) {
	emptiesX := make([]int, 0)
	emptiesY := make([]int, 0)
	for x := nx - 1; x <= nx+1; x++ {
		for y := ny - 1; y <= ny+1; y++ {
			if (nx-x)*(ny-y) != 0 { // restrict diagonals
				continue
			}
			if (x != nx || y != ny) && r.areCoordsValid(x, y) && r.elements[x][y].isEmpty() {
				emptiesX = append(emptiesX, x)
				emptiesY = append(emptiesY, y)
			}
		}
	}
	if len(emptiesX) == 0 {
		return -1, -1
	}
	index := r.rnd.Rand(len(emptiesX))
	return emptiesX[index], emptiesY[index]
}

func (r *LayoutMap) getRandomNonEmptyCellCoords(minEmptyCellsNear int) (int, int) {
	nonEmptiesX := make([]int, 0)
	nonEmptiesY := make([]int, 0)
	for x := 0; x < len(r.elements); x++ {
		for y := 0; y < len(r.elements[0]); y++ {
			// obstacles should not be counted as "non-empty"
			if !r.elements[x][y].isObstacle && !r.elements[x][y].isEmpty() && r.countEmptyCellsNear(x, y) >= minEmptyCellsNear {
				nonEmptiesX = append(nonEmptiesX, x)
				nonEmptiesY = append(nonEmptiesY, y)
			}
		}
	}
	if len(nonEmptiesX) == 0 {
		return -1, -1
	}
	index := r.rnd.Rand(len(nonEmptiesX))
	return nonEmptiesX[index], nonEmptiesY[index]
}

func (r *LayoutMap) getRandomPathCellCoords(desiredPathNum int, nodesAllowed bool) (int, int) { // desiredPathNum -1 means any path
	pathsX := make([]int, 0)
	pathsY := make([]int, 0)
	for x := 0; x < len(r.elements); x++ {
		for y := 0; y < len(r.elements[0]); y++ {
			if !r.elements[x][y].isPartOfAPath() {
				continue
			}
			if !nodesAllowed && r.elements[x][y].nodeInfo != nil { // don't take nodes unless allowed
				continue
			}
			if desiredPathNum > -1 && desiredPathNum != r.elements[x][y].pathInfo.pathNumber { // don't take cells of non-desired path numbers
				continue
			}
			pathsX = append(pathsX, x)
			pathsY = append(pathsY, y)
		}
	}
	if len(pathsX) == 0 {
		return -1, -1
	}
	index := r.rnd.Rand(len(pathsX))
	return pathsX[index], pathsY[index]
}

func (r *LayoutMap) areCoordsEmpty(x, y int) bool {
	return r.elements[x][y].isEmpty()
}

func (r *LayoutMap) isPathPresentAtCoords(x, y int) bool {
	return r.elements[x][y].isPartOfAPath()
}

func (r *LayoutMap) areCoordsEmptyOrPathOnly(x, y int) bool {
	return r.elements[x][y].isPathOrEmpty()
}

func (r *LayoutMap) countEmptyCellsNear(x, y int) int {
	count := 0
	w, h := r.GetSize()
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i*j != 0 || i == 0 && j == 0 {
				continue
			}
			if x+i < 0 || x+i >= w || y+j < 0 || y+j >= h {
				continue
			}
			if r.elements[x+i][y+j].isEmpty() {
				count++
			}
		}
	}
	return count
}

func (r *LayoutMap) getAnyOfCoordsOfNode(nodeName string) (int, int) {
	possibleCoords := make([][2]int, 0)
	for x := 0; x < len(r.elements); x++ {
		for y := 0; y < len(r.elements[0]); y++ {
			if r.elements[x][y].IsNode() && r.elements[x][y].nodeInfo.nodeName == nodeName {
				possibleCoords = append(possibleCoords, [2]int{x, y})
			}
		}
	}
	if len(possibleCoords) > 0 {
		rndIndex := r.rnd.Rand(len(possibleCoords))
		return possibleCoords[rndIndex][0], possibleCoords[rndIndex][1]
	}
	panic("getAnyOfCoordsOfNode failed with node " + nodeName)
	return -1, -1
}

func (r *LayoutMap) getAllCoordsOfNode(nodeName string) [][2]int {
	possibleCoords := make([][2]int, 0)
	for x := 0; x < len(r.elements); x++ {
		for y := 0; y < len(r.elements[0]); y++ {
			if r.elements[x][y].IsNode() && r.elements[x][y].nodeInfo.nodeName == nodeName {
				possibleCoords = append(possibleCoords, [2]int{x, y})
			}
		}
	}
	if len(possibleCoords) > 0 {
		return possibleCoords
	}
	panic("getAllCoordsOfNode failed with node " + nodeName)
}

func (l *LayoutMap) setAllNodeConnectionsLockedForPath(nodex, nodey, pathNum int, lockNum int) {
	e := l.GetElement(nodex, nodey)
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
			elem2 := l.GetElement(nodex+dirCoords[0], nodey+dirCoords[1])
			elem2.connections[opposingPath].IsLocked = true
			if elem2.connections[opposingPath].LockNum < lockNum {
				elem2.connections[opposingPath].LockNum = lockNum
			}
		}
	}
}

func (r *LayoutMap) areCoordsValid(x, y int) bool {
	w, h := r.GetSize()
	return x >= 0 && x < w && y >= 0 && y < h
}

// exported

func (r *LayoutMap) GetSize() (int, int) {
	return len(r.elements), len(r.elements[0])
}

func (r *LayoutMap) GetElement(x, y int) *element {
	return r.elements[x][y]
}

func (r *LayoutMap) getPassabilityMapForPathfinder(pathsArePassable bool) *[][]int {
	const (
		minRandomCostIncrease = -100
		maxRandomCostIncrease = 10000
	)
	layoutWidth, layoutHeight := r.GetSize()
	pmap := make([][]int, layoutWidth)
	for i := range pmap {
		pmap[i] = make([]int, layoutHeight)
	}

	for x := 0; x < layoutWidth; x++ {
		for y := 0; y < layoutHeight; y++ {
			if r.areCoordsEmpty(x, y) || pathsArePassable && r.areCoordsEmptyOrPathOnly(x, y) {
				pmap[x][y] = 1
				if r.isPathPresentAtCoords(x, y) {
					pmap[x][y] += maxRandomCostIncrease
				}
				// TODO: think how to better randomize path costs
				if r.randomizePaths {
					// lowering the "from" increases path randomness, but also makes the generator to fail more frequently
					// because it increases the probability for creating a non-existing path
					// "* 10" is to compensate the heuristics in the pathfinder
					pmap[x][y] += r.rnd.RandInRange(minRandomCostIncrease, maxRandomCostIncrease) * 10
				}
			} else {
				pmap[x][y] = -1
			}
		}
	}
	return &pmap
}
