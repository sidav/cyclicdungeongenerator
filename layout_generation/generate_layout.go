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
	var patternNumber int

	try := 0

	generationStart:
	for {
		try++
		layout.init(size, size)

		patternNumber = getRandomPatternNumber ()
		pattern := getPattern(patternNumber) // getPattern(1)

		for i := range pattern {
			// fmt.Printf("%d, ", i)
			success := execPatternStep(pattern[i])
			if !success {
				continue generationStart
			}
		}
		break
	}
	fmt.Printf("Generation finised, %d tries, final build pattern #%d \n", try, patternNumber)
	return &layout
}
