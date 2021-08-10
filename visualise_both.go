package main

import (
	cw "cyclicdungeongenerator/console_wrapper"
	layout_generation2 "cyclicdungeongenerator/generators/layout_generation"
	"cyclicdungeongenerator/random"
	"fmt"
)

type visBoth struct {
	currLayout       *layout_generation2.LayoutMap
	levelVis         tiledMapVisualiser
	layoutVis        layoutVisualiser
	currModeIsLayout bool // if false, curr mode is tiles

	flawsCritical, maxDesiredFlaws int
}

func (v *visBoth) do() {
	v.flawsCritical = 1000
	v.maxDesiredFlaws = 100
	rnd := random.FibRandom{}
	rnd.InitBySeed(-1)
	key := "none"
	desiredPatternNum := -1
	randomPaths := true
	parser := layout_generation2.PatternParser{}
	filenames := parser.ListPatternFilenamesInPath("patterns/")
	v.levelVis.roomW = 5
	v.levelVis.roomH = 5

	reGenerate := false
	pattNum := rnd.Rand(len(filenames))
	if desiredPatternNum != -1 {
		pattNum = desiredPatternNum
	}
	gen := layout_generation2.InitCyclicGenerator(randomPaths, W, H, -1)
	gen.TriesForPattern = v.flawsCritical
	pattern := parser.ParsePatternFile(filenames[pattNum], true)
	generatedMap, genRestarts := gen.GenerateLayout(pattern)

	for key != "ESCAPE" {
		cw.Clear_console()

		if reGenerate {
			pattNum = rnd.Rand(len(filenames))
			if desiredPatternNum != -1 {
				pattNum = desiredPatternNum
			}
			gen = layout_generation2.InitCyclicGenerator(randomPaths, W, H, -1)
			gen.TriesForPattern = v.flawsCritical
			pattern = parser.ParsePatternFile(filenames[pattNum], true)
			generatedMap, genRestarts = gen.GenerateLayout(pattern)
			reGenerate = false
		}

		if generatedMap == nil {
			key = ""
			cw.PutString(":(", 0, 0)
			cw.PutString(fmt.Sprintf("Generation failed even after %d restarts, pattern #%d", genRestarts, pattNum), 0, 1)
			cw.PutString("Press ENTER to generate again or ESCAPE to exit.", 0, 2)
			cw.Flush_console()
			for key != "ESCAPE" && key != "ENTER" {
				key = cw.ReadKey()
			}
			if key == "ENTER" {
				reGenerate = true
				continue
			} else {
				break
			}
		} else {
			if v.currModeIsLayout {
				v.layoutVis.putMap(generatedMap)
				v.layoutVis.putInfo(generatedMap, pattNum, desiredPatternNum, pattern.Filename, pattern.Name,
					genRestarts, v.maxDesiredFlaws, randomPaths)
			} else {
				v.levelVis.convertLayoutToLevelAndDraw(&rnd, generatedMap)
				v.levelVis.putInfo(generatedMap, pattNum, desiredPatternNum, pattern.Filename, pattern.Name,
					genRestarts, v.maxDesiredFlaws, gen.RandomizePath)
			}
		}
		cw.Flush_console()
		key = cw.ReadKey()
		switch key {
		case "r", "UP":
			randomPaths = !randomPaths
			reGenerate = true

		case "TAB", "m":
			v.currModeIsLayout = !v.currModeIsLayout
			reGenerate = false
			continue

		case "a":
			v.levelVis.roomW--

		case "d":
			v.levelVis.roomW++

		case "z":
			W--
			reGenerate = true

		case "x":
			W++
			reGenerate = true

		case "e":
			H--
			reGenerate = true

		case "c":
			H++
			reGenerate = true

		case "w":
			v.levelVis.roomH--

		case "s":
			v.levelVis.roomH++

		case "=", "+", "RIGHT":
			if desiredPatternNum < len(filenames)-1 {
				desiredPatternNum++
				reGenerate = true
			}

		case "-", "LEFT":
			if desiredPatternNum > -1 {
				desiredPatternNum--
				reGenerate = true
			}

		case "ENTER", "ESCAPE":
			reGenerate = true
		}
	}
}
