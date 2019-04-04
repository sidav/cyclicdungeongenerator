package layout_generation

var (
	size    = 5
	divisor = 3
	layout  = LayoutMap{}
)

func Generate(patternNumber int) (*LayoutMap, int) {
	const triesForPattern = 25

	if patternNumber == -1 {
		patternNumber = getRandomPatternNumber()
	}
	pattern := getPattern(patternNumber)

generationStart:
	for	generatorRestarts :=0; generatorRestarts <=triesForPattern; generatorRestarts++ {
		layout.init(size, size)

		for i := range pattern {
			success := execPatternStep(pattern[i])
			if !success {
				continue generationStart
			}
		}
		return &layout, generatorRestarts
	}
	return nil, triesForPattern
}
