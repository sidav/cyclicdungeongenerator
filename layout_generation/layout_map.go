package layout_generation

import (
	rnd "github.com/sidav/golibrl/random"
	"strconv"
)

type LayoutMap struct {
	// room map with connections or smth
	elements [][] *element
}

func (r *LayoutMap) init(sizex, sizey int) {
	r.elements = make([][]*element, size)
	for i := range r.elements {
		r.elements[i] = make([]*element, size)
	}
	for x := 0; x < sizex; x++ {
		for y := 0; y < sizey; y++ {
			r.elements[x][y] = &element{connections: map[string]*connection {"north": nil, "east": nil, "south": nil, "west": nil}}
		}
	}
}

func (r *LayoutMap) placeNodeAtCoords(x, y int, nodeName string) {
	r.elements[x][y].nodeInfo = &node_cell{nodeName: nodeName}
}

func (r *LayoutMap) placePathAtCoords(x, y int, pathNum int) {
	r.elements[x][y].pathInfo = &path_cell{pathNum}
}

func (r *LayoutMap) placeObstacleAtCoords(x, y int) {
	r.elements[x][y].isObstacle = true
}

func (r *LayoutMap) removeAllObstacles() {
	for x := 0; x < len(r.elements); x++ {
		for y := 0; y < len(r.elements[0]); y++ {
			r.elements[x][y].isObstacle = false
		}
	}
}

func (r *LayoutMap) getRandomPathCell(desiredPathNum int) (int, int) { // desiredPathNum -1 means any path
	x, y := rnd.Random(size), rnd.Random(size)
	const tries = 40
	try := 0
	for try < tries && (!r.elements[x][y].isPartOfAPath() || (desiredPathNum > -1 && desiredPathNum != r.elements[x][y].pathInfo.pathNumber)) {
		try++
		x, y = rnd.Random(size), rnd.Random(size)
	}
	return x, y
}

func (r *LayoutMap) getRandomPathCoordsAndRandomCellNearPath(pathNum int) (int, int, int, int) {
	const tries = 10
	for try := 0; try < tries; try++ {
		px, py := r.getRandomPathCell(pathNum)
		for try2 := 0; try2 < tries; try2++ {
			x, y := rnd.RandInRange(px-1, px+1), rnd.RandInRange(py-1, py+1)
			if x >= 0 && y >= 0 && x < len(r.elements) && y < len(r.elements[0]) && r.elements[x][y].isEmpty() {
				return px, py, x, y
			}
		}
	}
	return -1, -1, -1, -1
}

func (r *LayoutMap) getRandomEmptyCellCoords() (int, int) { // desiredPathNum -1 means any path
	const tries = 25
	for i := 0; i < tries; i++ {
		x, y := rnd.Random(size), rnd.Random(size)
		if r.elements[x][y].isEmpty() {
			return x, y
		}
	}
	return -1, -1
}

func (r *LayoutMap) areCoordsEmpty(x, y int) bool {
	return r.elements[x][y].isEmpty()
}

func (r *LayoutMap) getCoordsOfNode(nodeName string) (int, int) {
	for x := 0; x < len(r.elements); x++ {
		for y := 0; y < len(r.elements[0]); y++ {
			if r.elements[x][y].isNode() && r.elements[x][y].nodeInfo.nodeName == nodeName {
				return x, y
			}
		}
	}
	panic("getCoordsOfNode failed with node "+nodeName)
	return -1, -1
}

// exported

func (r *LayoutMap) GetSize() (int, int) {
	return len(r.elements), len(r.elements[0])
}

func (r *LayoutMap) getPassabilityMapForPathfinder() *[][]int {
	pmap := make([][]int, size)
	for i := range pmap {
		pmap[i] = make([]int, size)
	}

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			if layout.areCoordsEmpty(x, y) {
				pmap[x][y] = 1
			} else {
				pmap[x][y] = -1
			}
		}
	}
	return &pmap
}

// output TODO: remove

func (r *LayoutMap) GetCharOfElementAtCoords(x, y int) rune { // just for rendering, TODO: remove
	elem := r.elements[x][y]
	// rune := '?'
	if elem.isEmpty() {
		return '.'
	}
	if elem.isObstacle {
		return '#'
	}
	if elem.isNode() {
		return rune(elem.nodeInfo.nodeName[0])
	}
	if elem.isPartOfAPath() {
		number := elem.pathInfo.pathNumber
		return rune(strconv.Itoa(number)[0])
	}
	return '?'
}

func (r *LayoutMap) CellToCharArray(cellx, celly int) [][]rune {
	e := r.elements[cellx][celly]
	ca := make([][]rune, 5)
	for i := range (ca) {
		ca[i] = make([]rune, 5)
	}

	if e.isEmpty() {
		return ca
	}

	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			ca[x][y] = '#'
		}
	}
	if e.nodeInfo != nil {
		for x := 1; x < 4; x++ {
			for y := 1; y < 4; y++ {
				ca[x][y] = ' '
			}
		}
	}
	if e.pathInfo != nil {
		for x := 0; x < 5; x++ {
			for y := 0; y < 5; y++ {
				ca[x][y] = '.'
			}
		}
		ca[2][2] = rune(e.pathInfo.pathNumber - '0')
	}
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 2; y++ {
			if e.getConnectionByCoords(x, y) != nil {
				ca[2+x*2][2+y*2] = '+'
			}
		}
	}
	return ca
}
