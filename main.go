package main

import (
	"CyclicDungeonGenerator/layout_generation"
	cw "TCellConsoleWrapper"
	"fmt"
	rnd "github.com/sidav/golibrl/random"
)

func main() {
	rnd.Randomize()

	// layout_generation.Benchmark(-1)

	cw.Init_console()
	defer cw.Close_console()

	key := "none"
	desiredPatternNum := -1

	for key != "ESCAPE" {
		pattNum := rnd.Random(layout_generation.GetTotalPatternsNumber())
		if desiredPatternNum != -1 {
			pattNum = desiredPatternNum
		}
		generatedMap := layout_generation.Generate(pattNum)

		if generatedMap == nil {
			fmt.Printf("Failed.\n")
			return
		}

		cw.Clear_console()
		putMap(generatedMap)
		putMiniMapAndPatternNumber(generatedMap, pattNum, desiredPatternNum)
		cw.Flush_console()
keyread:
		for {
			key = cw.ReadKey()
			switch key {
			case "=":
				if desiredPatternNum < layout_generation.GetTotalPatternsNumber()-1 {
					desiredPatternNum++
				}
			break keyread
			case "-":
				if desiredPatternNum > -1 {
					desiredPatternNum--
				}
			break keyread
			case " ", "ESCAPE":
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


func putMiniMapAndPatternNumber(a *layout_generation.LayoutMap, pattNum, desiredPNum int) {
	sx, sy := a.GetSize()
	for y := 0; y < sy; y++ {
		for x := 0; x < sx; x++ {
			chr := a.GetCharOfElementAtCoords(x, y)
			setcolorForRune(chr)
			cw.PutChar(chr, x+sx*5 + 2, y)
		}
	}
	cw.SetFgColor(cw.BEIGE)
	cw.PutString(fmt.Sprintf("PATTERN SELECTED: #%d  ", desiredPNum), sx*5+2, sy+2)
	cw.PutString(fmt.Sprintf("PATTERN USED: #%d  ", pattNum), sx*5+2, sy+3)

}

func setcolorForRune(chr rune) {
	switch chr {
	case '1', '2', '3', '4', '5', '6':
		cw.SetFgColor(cw.DARK_CYAN)
	case '.':
		cw.SetFgColor(cw.BEIGE)
	case '+':
		cw.SetFgColor(cw.DARK_MAGENTA)
	case '#':
		cw.SetFgColor(cw.DARK_GRAY)
	default:
		cw.SetFgColor(cw.DARK_GREEN)
	}
}
