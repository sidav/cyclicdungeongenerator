package main

import (
	"CyclicDungeonGenerator/layout_generation"
	"CyclicDungeonGenerator/random"
	"fmt"
	cw "CyclicDungeonGenerator/console_wrapper"
)

func doLayoutVisualization() {
	rnd := random.FibRandom{}
	rnd.InitBySeed(-1)
	key := "none"
	desiredPatternNum := -1
	randomPaths := true
	parser := layout_generation.PatternParser{}
	filenames := parser.ListPatternFilenamesInPath("patterns/")

	for key != "ESCAPE" {
		cw.Clear_console()
		pattNum := rnd.Rand(len(filenames))
		if desiredPatternNum != -1 {
			pattNum = desiredPatternNum
		}
		gen := layout_generation.InitCyclicGenerator(randomPaths, W, H, -1)
		gen.TriesForPattern = 100
		pattern := parser.ParsePatternFile(filenames[pattNum])
		generatedMap, genRestarts := gen.GenerateLayout(pattern)

		if generatedMap == nil {
			cw.PutString(":(", 0, 0)
			cw.PutString(fmt.Sprintf("Generation failed even after %d restarts, pattern #%d", genRestarts, pattNum), 0, 1)
			cw.PutString("Press ENTER to generate again or ESCAPE to exit.", 0, 2)
			cw.Flush_console()
			for key != "ESCAPE" && key != "ENTER" {
				key = cw.ReadKey()
			}
			continue
		} else {
			putMap(generatedMap)
			putInfo(generatedMap, pattNum, desiredPatternNum, pattern.Filename, pattern.Name, genRestarts, randomPaths)
		}
		cw.Flush_console()
	keyread:
		for {
			key = cw.ReadKey()
			switch key {
			case "r", "UP":
				randomPaths = !randomPaths
				break keyread
			case "=", "+", "RIGHT":
				if desiredPatternNum < len(filenames)-1 {
					desiredPatternNum++
				}
				break keyread
			case "-", "LEFT":
				if desiredPatternNum > -1 {
					desiredPatternNum--
				}
				break keyread
			case " ", "TAB", "ESCAPE":
				break keyread
			}
		}
	}
}

func putCharArray(x, y int, c *[][]rune) {
	for i := 0; i < len(*c); i++ {
		for j := 0; j < len((*c)[0]); j++ {
			setcolorForRune((*c)[i][j])
			cw.PutChar((*c)[i][j], x+i, y+j)
		}
	}
}

func putMap(a *layout_generation.LayoutMap) {
	putCharArray(0, 0, a.WholeMapToCharArray())
}

func putInfo(a *layout_generation.LayoutMap, pattNum, desiredPNum int, fName, pName string, restarts int, rand bool) {
	sx, sy := a.GetSize()
	for y := 0; y < sy; y++ {
		for x := 0; x < sx; x++ {
			chr := a.GetCharOfElementAtCoords(x, y)
			setcolorForRune(chr)
			cw.PutChar(chr, x+sx*5+2, y)
		}
	}
	cw.SetFgColor(cw.BEIGE)
	cw.PutString(fmt.Sprintf("PATTERN SELECTED: #%d  ", desiredPNum), sx*5+2, sy+2)
	cw.PutString(fmt.Sprintf("PATTERN USED: #%d  ", pattNum), sx*5+2, sy+3)
	cw.PutString(fmt.Sprintf("FILE: %s  ", fName), sx*5+2, sy+4)
	cw.PutString(fmt.Sprintf("NAME: %s  ", pName), sx*5+2, sy+5)
	cw.PutString(fmt.Sprintf("Gen restarts: %d", restarts), sx*5+2, sy+6)
	if rand {
		cw.PutString("Random paths", sx*5+2, sy+7)
	} else {
		cw.PutString("Shortest paths", sx*5+2, sy+7)
	}
}

func setcolorForRune(chr rune) {
	cw.SetBgColor(cw.BLACK)
	switch chr {
	case '1', '2', '3', '4', '5', '6':
		cw.SetFgColor(cw.DARK_CYAN)
	case '.':
		cw.SetFgColor(cw.BEIGE)
	case '~':
		cw.SetFgColor(cw.BLUE)
	case '+':
		cw.SetFgColor(cw.DARK_MAGENTA)
	case '#':
		cw.SetColor(cw.DARK_GRAY, cw.DARK_GRAY)
	case '%':
		cw.SetColor(cw.BLACK, cw.RED)
	case '=':
		cw.SetColor(cw.BLACK, cw.DARK_BLUE)
	default:
		cw.SetFgColor(cw.DARK_GREEN)
	}
}
