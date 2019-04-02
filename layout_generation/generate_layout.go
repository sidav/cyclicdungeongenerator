package layout_generation

import (
	"fmt"
	rnd "github.com/sidav/golibrl/random"
)

var (
	size    = 5
	divisor = 3
	layout  = LayoutMap{}
)

func Generate() *LayoutMap {
	rnd.Randomize()

	const triesForPattern = 10

	patternNumber := getRandomPatternNumber()
	pattern := getPattern(patternNumber)

generationStart:
	for	patternTry:=1;patternTry<=triesForPattern; patternTry++ {
		layout.init(size, size)

		for i := range pattern {
			success := execPatternStep(pattern[i])
			if !success {
				continue generationStart
			}
		}
		fmt.Printf("Generation finised, %d tries, final build pattern #%d \n", patternTry, patternNumber)
		return &layout
	}
	fmt.Printf("Generation failed for pattern #%d after %d tries\n", patternNumber, triesForPattern)
	return nil
}
