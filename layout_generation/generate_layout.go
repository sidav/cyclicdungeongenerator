package layout_generation

import (
	"fmt"
	rnd "github.com/sidav/golibrl/random"
)

var (
	size        = 5
	divisor  = 3
	layout = LayoutMap{}
)

const (
)

func Generate() *LayoutMap {
	rnd.Randomize()

	generationStart:
	for {

		layout.init(size, size)

		pattern := getPattern(0)

		for i := range pattern {
			fmt.Printf("%d, ", i)
			success := execPatternStep(pattern[i])
			if !success {
				continue generationStart
			}
		}
		break
	}
	return &layout
}
