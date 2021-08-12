package layout_generation

import (
	"cyclicdungeongenerator/random"
)

type CyclicGenerator struct {
	RandomizePath             bool
	layoutWidth, layoutHeight int
	layout                    LayoutMap
	rnd                       random.FibRandom
	TriesForPattern           int
}

func InitCyclicGenerator(randomizePath bool, layoutWidth, layoutHeight int, seed int) *CyclicGenerator {
	gen := &CyclicGenerator{
		RandomizePath:   randomizePath,
		layoutWidth:     layoutWidth,
		layoutHeight:    layoutHeight,
		layout:          LayoutMap{},
		rnd:             random.FibRandom{},
		TriesForPattern: 100,
	}
	gen.rnd.InitBySeed(seed)
	return gen
}

func (cg *CyclicGenerator) GenerateLayout(pattern *pattern) (*LayoutMap, int) {

generationStart:
	for generatorRestarts := 0; generatorRestarts <= cg.TriesForPattern; generatorRestarts++ {
		cg.layout.init(cg.layoutWidth, cg.layoutHeight, &cg.rnd, cg.RandomizePath)

		for i := range pattern.instructions {
			success := pattern.instructions[i].execPatternStep(&cg.layout)
			if !success {
				continue generationStart
			}
		}
		cg.layout.randomizeTagLocationsPerNode()
		return &cg.layout, generatorRestarts
	}
	return nil, cg.TriesForPattern
}
