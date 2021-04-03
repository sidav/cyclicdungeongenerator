package layout_generation

import (
	"fmt"
	"strings"
)

type Benchmark struct {
	BenchLoopsForPattern            int
	TriesForPattern                 int
	CheckRandomPaths                bool
	CheckShortestPaths              bool
	TestUniquity                    bool
	GenerateAndConsiderGarbageNodes bool
}

func (b *Benchmark) Benchmark(patternNum int) {
	if !b.CheckRandomPaths {
		fmt.Println("Not testing random paths.")
	}
	if !b.CheckShortestPaths {
		fmt.Println("Not testing shortest paths.")
	}
	if b.CheckRandomPaths || b.CheckShortestPaths {
		if patternNum == -1 {
			fmt.Printf("\rBENCHMARK FOR ALL PATTERNS:\n")
			for i := 0; i < GetTotalPatternsNumber(); i++ {
				b.startBench(i)
			}
		} else {
			fmt.Printf("\rBENCHMARKING PATTERN %d:\n", patternNum)
			b.startBench(patternNum)
		}
		fmt.Printf("Benchmark finished. Press Enter. \n")
	} else {
		fmt.Printf("Nothing to benchmark. Press Enter. \n")
	}
	var input string
	fmt.Scanln(&input)
}

func (b *Benchmark) startBench(patternNum int) {
	if b.CheckRandomPaths {
		fmt.Printf("BENCHMARKING #%d, RANDOM PATHS: \n", patternNum)
		RandomizePath = true
		b.benchmarkPattern(patternNum, b.TestUniquity, b.GenerateAndConsiderGarbageNodes)
	}
	if b.CheckShortestPaths {
		fmt.Printf("BENCHMARKING #%d, SHORTEST PATHS: \n", patternNum)
		RandomizePath = false
		b.benchmarkPattern(patternNum, b.TestUniquity, b.GenerateAndConsiderGarbageNodes)
	}
}

func (b *Benchmark) getCharmapAndTriesAndSuccessForGeneration(patternNumber int, countGarbageNodes bool) (*[][]rune, int, bool, *[]int) {

	if patternNumber == -1 {
		patternNumber = getRandomPatternNumber()
	}
	pattern := getPattern(patternNumber)
	flawsPerStep := make([]int, len(pattern))

generationStart:
	for patternTry := 0; patternTry <= b.TriesForPattern; patternTry++ {
		layout.init(layoutWidth, layoutHeight)

		for i := range pattern {
			if !countGarbageNodes {
				if pattern[i].actionType == ACTION_PLACE_RANDOM_CONNECTED_NODES ||
					pattern[i].actionType == ACTION_FILL_WITH_RANDOM_CONNECTED_NODES {
					continue // don't count those random unneccessary nodes.
				}
			}
			success := execPatternStep(pattern[i])
			if !success {
				flawsPerStep[i]++
				continue generationStart
			}
		}
		return layout.WholeMapToCharArray(), patternTry, true, &flawsPerStep
	}
	return nil, b.TriesForPattern, false, &flawsPerStep
}

func (b *Benchmark) benchmarkPattern(patternNum int, testUniquity bool, countGarbageNodes bool) {
	generatedMaps := make([]*[][]rune, 0)
	maxSteps := 0
	minSteps := 99999999
	stepsSum := 0
	fails := 0
	repeats := 0
	flawsPerStep := make([]int, len(getPattern(patternNum)))
	for loopNum := 0; loopNum < b.BenchLoopsForPattern; loopNum++ {
		progressBarCLI(fmt.Sprintf("Progress "), loopNum+1, b.BenchLoopsForPattern+1, 15)
		cMap, tries, success, flawsPerGeneration := b.getCharmapAndTriesAndSuccessForGeneration(patternNum, countGarbageNodes)
		if testUniquity && cMap != nil {
			if !b.isCharmapAlreadyInArray(cMap, &generatedMaps) {
				generatedMaps = append(generatedMaps, cMap)
			} else {
				repeats++
			}
		}
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
		for i := 0; i < len(flawsPerStep); i++ {
			flawsPerStep[i] += (*flawsPerGeneration)[i]
		}
	}

	fmt.Printf("Pattern #%d, min flaws %d, max flaws %d, mean flaws count %.2f, %d failed attempts (%.2f%%)\n", patternNum,
		minSteps, maxSteps, float64(stepsSum)/float64(b.BenchLoopsForPattern), fails,
		100.0*float64(fails)/(float64(b.BenchLoopsForPattern)))
	fmt.Print("Flaws per step: \n")
	flawsArrString := ""
	for i := 0; i < len(flawsPerStep); i++ {
		flawsArrString += fmt.Sprintf("%d: %d;  ", i, flawsPerStep[i])
	}
	fmt.Print(flawsArrString + "\n")

	if testUniquity {
		fmt.Printf("There was %d unique maps and %d repeats, repeats consist %.2f%% of total maps generated).\n\n",
			len(generatedMaps), repeats, 100.0*float64(repeats)/float64(repeats+len(generatedMaps)))
	} else {
		fmt.Printf("Uniquity test was not performed as set by TestUniquity flag. \n")
	}
	fmt.Print("\n")
}

func (b *Benchmark) isCharmapAlreadyInArray(c *[][]rune, arr *[]*[][]rune) bool {
	for i := 0; i < len(*arr); i++ {
		if b.areTwoCharArraysEqual(c, (*arr)[i]) {
			return true
		}
	}
	return false
}

func (b *Benchmark) areTwoCharArraysEqual(c1, c2 *[][]rune) bool {
	if len(*c1) != len(*c2) {
		return false
	}
	for i := 0; i < len(*c1); i++ {
		for j := 0; j < len((*c1)[0]); j++ {
			if (*c1)[i][j] != (*c2)[i][j] {
				return false
			}
		}
	}
	return true
}

func progressBarCLI(title string, value, endvalue, bar_length int) { // because I can
	endvalue -= 1
	percent := float64(value) / float64(endvalue)
	arrow := ">"
	for i := 0; i < int(percent*float64(bar_length)); i++ {
		arrow = "-" + arrow
	}
	spaces := strings.Repeat(" ", bar_length-len(arrow)+1)
	percent_with_dec := fmt.Sprintf("%.1f", percent*100.0)
	fmt.Printf("\r%s [%s%s]%s%% (%d out of %d)", title, arrow, spaces, percent_with_dec, value, endvalue)
	if value == endvalue {
		fmt.Printf("\n")
	}
}
