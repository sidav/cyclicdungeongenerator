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

func Generate(patternNumber int) *LayoutMap {
	rnd.Randomize()

	const triesForPattern = 10

	if patternNumber == -1 {
		patternNumber = getRandomPatternNumber()
	}
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

// benchmarking below

func getTriesAndSuccessForGeneration(patternNumber int) (int, bool) {
	rnd.Randomize()

	const triesForPattern = 1000

	if patternNumber == -1 {
		patternNumber = getRandomPatternNumber()
	}
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
		return patternTry, true
	}
	// fmt.Printf("Generation failed for pattern #%d after %d tries\n", patternNumber, triesForPattern)
	return triesForPattern, false
}

func Benchmark(patternNum int) {
	const benchLoopsForPattern = 100000
	maxSteps := 0
	minSteps := 99999999
	stepsSum := 0
	fails := 0
	for loopNum := 0; loopNum < benchLoopsForPattern; loopNum++ {
		tries , success := getTriesAndSuccessForGeneration(patternNum)
		stepsSum += tries
		if maxSteps < tries {
			maxSteps = tries
		}
		if minSteps > tries {
			minSteps = tries
		}
		if !success {
			fails++
		}
	}

	fmt.Printf("Pattern #%d, min tries %d, max tries %d, mean tries number %f, %d failed attempts\n", patternNum,
		minSteps, maxSteps, float64(stepsSum)/float64(benchLoopsForPattern), fails)
}
