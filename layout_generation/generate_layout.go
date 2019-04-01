package layout_generation

import rnd "github.com/sidav/golibrl/random"

var (
	size        = 5
	divisor  = 3
	layout = LayoutMap{}
)

const (
)

func Generate() *LayoutMap {
	rnd.Randomize()
	layout.init(size, size)

	pattern := getPattern(0)

	for i := range pattern {
		execPatternStep(pattern[i])
	}

	return &layout
}

func GenerateDeprecated() *LayoutMap {
	rnd.Randomize()

	layout.init(size, size)

	layout.removeAllObstacles()

	// add node to path
	//nx, ny := layout.getRandomPathCell(-1)
	//layout.placeNodeAtCoords(nx, ny, 'N')

	return &layout
}
//
//func addNodeToPathAtRandom() {
//	nx, ny := layout.getRandomPathCell(-1)
//	layout.placeNodeAtCoords(nx, ny, 'N')
//}
