package layout_generation

import rnd "github.com/sidav/goLibRL/random"
import "github.com/sidav/golibrl/astar"

var (
	size        = 10
	divisor  = 3
	layout = LayoutMap{}
)

const (
	null_rune     = rune(0)
	temp_obstacle = ';'
)

func Generate() *LayoutMap {
	rnd.Randomize()

	layout.init(size, size)

	// place beginning randomly
	fx, fy := rnd.Random(size/divisor), rnd.Random(size/divisor)
	layout.placeNodeAtCoords(fx, fy, 'S')

	// place end randomly
	tx, ty := size-1-rnd.Random(size/divisor), size-1-rnd.Random(size/divisor)
	layout.placeNodeAtCoords(tx, ty, 'F')

	// place big obstacle in center and some random obstacles for path to be less straight
	for i := 0; i < size/3+1; i++ {
		for j := 0; j < size/3+1; j++ {
			layout.placeObstacleAtCoords(size/3+i, size/3+j)
		}
	}

	rnd_obstcls_count := rnd.RandInRange(size/2, size/2+size/4)
	placeTempObstacles(rnd_obstcls_count)

	// draw the path itself
	findAndDrawPathFromTo(fx, fy, tx, ty, 1)

	// draw the second path
	findAndDrawPathFromTo(fx, fy, tx, ty, 2)

	layout.removeAllObstacles()

	// add node to path

	nx, ny := layout.getRandomPathCell(-1)
	layout.placeNodeAtCoords(nx, ny, 'N')

	placeTempObstacles(5)
	findAndDrawPathFromTo(nx, ny, tx, ty, 3)

	return &layout
}

func addNodeToPathAtRandom() {
	nx, ny := layout.getRandomPathCell(-1)
	layout.placeNodeAtCoords(nx, ny, 'N')
}

func findAndDrawPathFromTo(fx, fy, tx, ty, pathNumber int) {
	pmap := getPassabilityMapForPathfinder()
	path := astar.FindPath(pmap, fx, fy, tx, ty, false, true)
	if path == nil {
		return
	}
	for path.Child != nil {
		path = path.Child
		x, y := path.GetCoords()
		layout.placePathAtCoords(x, y, pathNumber)
	}
}

func getPassabilityMapForPathfinder() *[][]int {
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

func placeTempObstacles(count int) {
	for i := 0; i < count; i++ {
		x, y := layout.getRandomEmptyCellCoords()
		layout.placeObstacleAtCoords(x, y)
	}
}
