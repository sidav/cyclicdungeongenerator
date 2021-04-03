package layout_generation

import (
	"CyclicDungeonGenerator/random"
)

type CyclicGenerator struct {
	RandomizePath             bool
	layoutWidth, layoutHeight int
	layout                    LayoutMap
	rnd                       random.FibRandom
}

func InitCyclicGenerator(randomizePath bool, layoutWidth, layoutHeight int, seed int) *CyclicGenerator {
	gen := &CyclicGenerator{
		RandomizePath: randomizePath,
		layoutWidth:   layoutWidth,
		layoutHeight:  layoutHeight,
		layout:        LayoutMap{},
		rnd:           random.FibRandom{},
	}
	gen.rnd.InitBySeed(seed)
	return gen
}

func (cg *CyclicGenerator) GenerateLayout(patternNumber int) (*LayoutMap, int) {
	const triesForPattern = 25

	if patternNumber == -1 {
		patternNumber = getRandomPatternNumber(&cg.rnd)
	}
	pattern := getPattern(patternNumber)

	// TODO: DELETE
	pp := PatternParser{}
	pattern = pp.ParsePatternFile("example_pattern.ptn")
	// TODO: DELETE ABOVE THIS LINE

generationStart:
	for generatorRestarts := 0; generatorRestarts <= triesForPattern; generatorRestarts++ {
		cg.layout.init(cg.layoutWidth, cg.layoutHeight, &cg.rnd, cg.RandomizePath)

		for i := range pattern.instructions {
			success := pattern.instructions[i].execPatternStep(&cg.layout)
			if !success {
				continue generationStart
			}
		}
		return &cg.layout, generatorRestarts
	}
	return nil, triesForPattern
}
