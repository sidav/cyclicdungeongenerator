package layout_generation

import (
	"CyclicDungeonGenerator/random"
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

	layout                    LayoutMap
	randomizePaths            bool
	LayoutWidth, LayoutHeight int
	rnd                       random.FibRandom
	parser                    PatternParser
}

func (b *Benchmark) Benchmark(patternFilename string) {
	b.rnd = random.FibRandom{}
	b.rnd.InitBySeed(-1)
	if !b.CheckRandomPaths {
		fmt.Println("Not testing random paths.")
	}
	if !b.CheckShortestPaths {
		fmt.Println("Not testing shortest paths.")
	}
	if b.CheckRandomPaths || b.CheckShortestPaths {
		if !strings.Contains(patternFilename, ".ptn") {
			fmt.Printf("\rBENCHMARK FOR ALL PATTERNS:\n")
			filenames := b.parser.ListPatternFilenamesInPath(patternFilename)
			for i := range filenames {
					b.startBench(filenames[i])
			}
		} else {
			fmt.Printf("\rBENCHMARKING PATTERN %s:\n", patternFilename)
			b.startBench(patternFilename)
		}
		fmt.Printf("Benchmark finished. Press Enter. \n")
	} else {
		fmt.Printf("Nothing to benchmark. Press Enter. \n")
	}
	var input string
	fmt.Scanln(&input)
}

func (b *Benchmark) startBench(patternFilename string) {
	pattern := b.parser.ParsePatternFile(patternFilename)
	if b.CheckRandomPaths {
		fmt.Printf("BENCHMARKING %s, RANDOM PATHS: \n", pattern.Filename)
		b.randomizePaths = true
		b.benchmarkPattern(pattern, b.TestUniquity, b.GenerateAndConsiderGarbageNodes)
	}
	if b.CheckShortestPaths {
		fmt.Printf("BENCHMARKING %s, SHORTEST PATHS: \n", pattern.Filename)
		b.randomizePaths = false
		b.benchmarkPattern(pattern, b.TestUniquity, b.GenerateAndConsiderGarbageNodes)
	}
}

func (b *Benchmark) getCharmapAndTriesAndSuccessForGeneration(pattern *pattern, countGarbageNodes bool) (*[][]rune, int, bool, *[]int) {
	flawsPerStep := make([]int, len(pattern.instructions))

generationStart:
	for patternTry := 0; patternTry <= b.TriesForPattern; patternTry++ {
		b.layout.init(b.LayoutWidth, b.LayoutHeight, &b.rnd, b.randomizePaths)

		for i := range pattern.instructions {
			if !countGarbageNodes {
				if pattern.instructions[i].actionType == ACTION_PLACE_RANDOM_CONNECTED_NODES ||
					pattern.instructions[i].actionType == ACTION_FILL_WITH_RANDOM_CONNECTED_NODES {
					continue // don't count those random unneccessary nodes.
				}
			}
			success := pattern.instructions[i].execPatternStep(&b.layout)
			if !success {
				flawsPerStep[i]++
				continue generationStart
			}
		}
		return b.layout.WholeMapToCharArray(), patternTry, true, &flawsPerStep
	}
	return nil, b.TriesForPattern, false, &flawsPerStep
}

func (b *Benchmark) benchmarkPattern(pattern *pattern, testUniquity bool, countGarbageNodes bool) {
	generatedMaps := make([]*[][]rune, 0)
	maxSteps := 0
	minSteps := 99999999
	stepsSum := 0
	fails := 0
	repeats := 0
	flawsPerStep := make([]int, len(pattern.instructions))
	for loopNum := 0; loopNum < b.BenchLoopsForPattern; loopNum++ {
		progressBarCLI(fmt.Sprintf("Progress "), loopNum+1, b.BenchLoopsForPattern+1, 15)
		cMap, tries, success, flawsPerGeneration := b.getCharmapAndTriesAndSuccessForGeneration(pattern, countGarbageNodes)
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

	fmt.Printf("Pattern #%s, min flaws %d, max flaws %d, mean flaws count %.2f, %d failed attempts (%.2f%%)\n", pattern.Name,
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
