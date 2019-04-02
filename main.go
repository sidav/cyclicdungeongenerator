package main

import (
	"CyclicDungeonGenerator/layout_generation"
	cw "TCellConsoleWrapper"
	rnd "github.com/sidav/golibrl/random"
)

func main() {
	rnd.Randomize()

	generatedMap := layout_generation.Generate()

	cw.Init_console()
	defer cw.Close_console()

	putMap(generatedMap)
	putMiniMap(generatedMap)
	cw.Flush_console()
	cw.ReadKey()
}

func putCharArray(x, y int, c [][]rune) {
	for i := 0; i < len(c); i++ {
		for j := 0; j < len(c[0]); j++ {
			setcolorForRune(c[i][j])
			cw.PutChar(c[i][j], x+i, y+j)
		}
	}
}

func putMap(a *layout_generation.LayoutMap) {
	sx, sy := a.GetSize()
	for y := 0; y < sy; y++ {
		for x := 0; x < sx; x++ {
			ca := a.CellToCharArray(x, y)
			putCharArray(x*5, y*5, ca)
		}
	}
}


func putMiniMap(a *layout_generation.LayoutMap) {
	sx, sy := a.GetSize()
	for y := 0; y < sy; y++ {
		for x := 0; x < sx; x++ {
			chr := a.GetCharOfElementAtCoords(x, y)
			setcolorForRune(chr)
			cw.PutChar(chr, x+sx*5 + 2, y)
		}
	}
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
