package cyc_dung_gen

import rnd "github.com/sidav/goLibRL/random"
import "github.com/sidav/golibrl/astar"

var (
	size    = 10
	divisor = 3
	lay     [][]rune
)

const (
	null_rune     = rune(0)
	temp_obstacle = ';'
)

func Generate() *[][]rune {
	rnd.Randomize()

	lay = make([][]rune, size)
	for i := range lay {
		lay[i] = make([]rune, size)
	}
	// place beginning randomly
	fx, fy := rnd.Random(size/divisor), rnd.Random(size/divisor)
	lay[fx][fy] = 'S'

	// place end randomly
	tx, ty := size-1-rnd.Random(size/divisor), size-1-rnd.Random(size/divisor)
	lay[tx][ty] = 'F'

	// place big obstacle in center and some random obstacles for path to be less straight
	for i := 0; i < size/3+1; i++ {
		for j := 0; j < size/3+1; j++ {
			lay[size/3+i][size/3+j] = temp_obstacle
		}
	}
	rnd_obstcls_count := rnd.RandInRange(size/2, size)
	for i := 0; i < rnd_obstcls_count; i++ {
		x, y := getRandomCoordsWithChar(null_rune)
		lay[x][y] = temp_obstacle
	}

	// draw the path itself
	findAndDrawPathFromTo(fx, fy, tx, ty)

	// draw the second path
	findAndDrawPathFromTo(fx, fy, tx, ty)

	clearTemps()
	// shit dots on empty cells
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			if lay[x][y] == null_rune {
				lay[x][y] = '.'
			}
		}
	}

	return &lay
}

func findAndDrawPathFromTo(fx, fy, tx, ty int) {
	pmap := getPassabilityMapForPathfinder()
	path := astar.FindPath(pmap, fx, fy, tx, ty, false, true)
	if path == nil {
		return
	}
	for path.Child != nil {
		path = path.Child
		x, y := path.GetCoords()
		lay[x][y] = '*'
	}
}

func getPassabilityMapForPathfinder() *[][]int {
	pmap := make([][]int, size)
	for i := range pmap {
		pmap[i] = make([]int, size)
	}

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			if lay[x][y] == null_rune {
				pmap[x][y] = 1
			} else {
				pmap[x][y] = -1
			}
		}
	}
	return &pmap
}

func getRandomCoordsWithChar(chr rune) (int, int) {
	x, y := rnd.Random(size), rnd.Random(size)
	for lay[x][y] != chr {
		x, y = rnd.Random(size), rnd.Random(size)
	}
	return x, y
}

func clearTemps() {
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			if lay[x][y] == temp_obstacle {
				lay[x][y] = null_rune
			}
		}
	}
}
