package layout_generation

var (
	layoutWidth  = 5
	layoutHeight = 5
	layout       = LayoutMap{}
)

func Generate(patternNumber int) (*LayoutMap, int) {
	const triesForPattern = 25

	if patternNumber == -1 {
		patternNumber = getRandomPatternNumber()
	}
	pattern := getPattern(patternNumber)

generationStart:
	for generatorRestarts := 0; generatorRestarts <= triesForPattern; generatorRestarts++ {
		layout.init(layoutWidth, layoutHeight)

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
