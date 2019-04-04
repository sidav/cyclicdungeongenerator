package layout_generation

import (
	"fmt"
	"strings"
)

const benchLoopsForPattern = 10000

func Benchmark(patternNum int, testUniquity bool) {
	if patternNum == -1 {
		fmt.Printf("\rBENCHMARK FOR ALL PATTERNS:\n")
		for i := 0; i < GetTotalPatternsNumber(); i++ {
			benchmarkPattern(i, testUniquity)
		}
	} else {
		fmt.Printf("\rBENCHMARKING PATTERN %d:\n", patternNum)
		benchmarkPattern(patternNum, testUniquity)
	}
	fmt.Printf("Benchmark finished. Press Enter. \n")
	var input string
	fmt.Scanln(&input)
}

func getCharmapAndTriesAndSuccessForGeneration(patternNumber int) (*[][]rune, int, bool, *[]int) {
	const triesForPattern = 1000

	if patternNumber == -1 {
		patternNumber = getRandomPatternNumber()
	}
	pattern := getPattern(patternNumber)
	flawsPerStep := make([]int, len(pattern))

generationStart:
	for	patternTry:=0;patternTry<=triesForPattern; patternTry++ {
		layout.init(size, size)

		for i := range pattern {
			success := execPatternStep(pattern[i])
			if !success {
				flawsPerStep[i]++
				continue generationStart
			}
		}
		return layout.WholeMapToCharArray(), patternTry, true, &flawsPerStep
	}
	return nil, triesForPattern, false, &flawsPerStep
}

func benchmarkPattern(patternNum int, testUniquity bool) {
	generatedMaps := make([]*[][]rune, 0)
	maxSteps := 0
	minSteps := 99999999
	stepsSum := 0
	fails := 0
	repeats := 0
	flawsPerStep := make([]int, len(getPattern(patternNum)))
	for loopNum := 0; loopNum < benchLoopsForPattern; loopNum++ {
		progressBarCLI(fmt.Sprintf("Benchmarking pattern #%d", patternNum), loopNum, benchLoopsForPattern, 20)
		cMap, tries , success, flawsPerGeneration := getCharmapAndTriesAndSuccessForGeneration(patternNum)
		if testUniquity {
			if !isCharmapAlreadyInArray(cMap, &generatedMaps) {
				generatedMaps = append(generatedMaps, cMap)
			} else {
				repeats ++
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
		for i:=0;i<len(flawsPerStep);i++{
			flawsPerStep[i] += (*flawsPerGeneration)[i]
		}
	}

	fmt.Printf("Pattern #%d, min flaws %d, max flaws %d, mean flaws count %f, %d failed attempts\n", patternNum,
		minSteps, maxSteps, float64(stepsSum)/float64(benchLoopsForPattern), fails)
	fmt.Print("Flaws per step: \n")
	flawsArrString := ""
	for i:=0;i<len(flawsPerStep);i++{
		flawsArrString += fmt.Sprintf("%d: %d;  ", i, flawsPerStep[i])
	}
	fmt.Print(flawsArrString + "\n")

	if testUniquity {
		fmt.Printf("There was %d unique maps and %d repeats, repeats consist %.2f%% of total maps generated).\n\n",
			len(generatedMaps), repeats, 100.0*float64(repeats)/float64(repeats+len(generatedMaps)))
	} else {
		fmt.Printf("Uniquity test was not performed as set by testUniquity flag. \n")
	}
	fmt.Print("\n")
}

func isCharmapAlreadyInArray(c *[][]rune, arr *[]*[][]rune) bool {
	for i := 0;i<len(*arr);i++{
		if areTwoCharArraysEqual(c, (*arr)[i]) {
			return true
		}
	}
	return false
}

func areTwoCharArraysEqual(c1, c2 *[][]rune) bool {
	if len(*c1) != len(*c2) {
		return false
	}
	for i:=0;i<len(*c1);i++{
		for j:=0;j<len((*c1)[0]);j++ {
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
	for i:=0; i < int(percent * float64(bar_length)); i++ {
		arrow = "-" + arrow
	}
	spaces := strings.Repeat(" ", bar_length - len(arrow) + 1)
	percent_with_dec := fmt.Sprintf("%.2f", percent*100.0)
	fmt.Printf("\r%s [%s%s]%s%% (%d out of %d)", title, arrow, spaces, percent_with_dec, value, endvalue)
	if value == endvalue {
		fmt.Printf("\n")
	}
}
