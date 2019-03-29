package cyc_dung_gen

import rnd "github.com/sidav/goLibRL/random"

var layout *rmap

func Generate() {
	size := 10
	rnd.Randomize()

	lay := make([][]rune, size)
	for i := range lay {
		lay[i] = make([]rune, size)
	}
	// place beginning randomly
	x, y := 0 + rnd.Random(size/2), 0 + rnd.Random(size/2)
	lay[x][y] = 'S'

}

func getPassabilityMapForPathfinder() *[][]bool {

}
