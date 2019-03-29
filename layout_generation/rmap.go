package layout_generation

import (
	rnd "github.com/sidav/goLibRL/random"
	"strconv"
)

type LayoutMap struct {// room map with connections or smth
	elements [][] *element
}

func (r *LayoutMap) init(sizex, sizey int) {
	r.elements = make([][]*element, size)
	for i := range r.elements {
		r.elements[i] = make([]*element, size)
	}
	for x:=0;x<sizex;x++{
		for y :=0; y <sizey; y++{
			r.elements[x][y] = &element{}
		}
	}
}

func (r *LayoutMap) placeNodeAtCoords(x, y int, nodeRune rune) {
	r.elements[x][y].nodeInfo = &node_cell{nodeChar: nodeRune}
}

func (r *LayoutMap) placePathAtCoords(x, y int, pathNum int) {
	r.elements[x][y].pathInfo = &path_cell{pathNum}
}

func (r *LayoutMap) placeObstacleAtCoords(x, y int) {
	r.elements[x][y].isObstacle = true
}

func (r *LayoutMap) removeAllObstacles() {
	for x:=0;x<len(r.elements);x++{
		for y :=0; y <len(r.elements[0]); y++{
			r.elements[x][y].isObstacle = false
		}
	}
}

func (r *LayoutMap) getRandomPathCell(desiredPathNum int) (int, int) { // desiredPathNum -1 means any path
	x, y := rnd.Random(size), rnd.Random(size)
	for !r.elements[x][y].isPartOfAPath() || (desiredPathNum > -1 && desiredPathNum != r.elements[x][y].pathInfo.pathNumber)  {
		x, y = rnd.Random(size), rnd.Random(size)
	}
	return x, y
}

func (r *LayoutMap) getRandomEmptyCellCoords() (int, int) { // desiredPathNum -1 means any path
	x, y := rnd.Random(size), rnd.Random(size)
	for !r.elements[x][y].isEmpty()  {
		x, y = rnd.Random(size), rnd.Random(size)
	}
	return x, y
}

func (r *LayoutMap) areCoordsEmpty(x, y int) bool {
	return r.elements[x][y].isEmpty()
}

// exported

func (r *LayoutMap) GetSize() (int, int) {
	return len(r.elements), len(r.elements[0])
}

func (r *LayoutMap) GetCharOfElementAtCoords(x, y int) rune { // just for rendering, TODO: remove
	elem := r.elements[x][y]
	// rune := '?'
	if elem.isEmpty() {
		return '.'
	}
	if elem.isObstacle{
		return '#'
	}
	if elem.isNode() {
		return elem.nodeInfo.nodeChar
	}
	if elem.isPartOfAPath() {
		number := elem.pathInfo.pathNumber
		return rune(strconv.Itoa(number)[0])
	}
	return '?'
}
