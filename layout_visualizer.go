package main

import (
	cw "CyclicDungeonGenerator/console_wrapper"
	"CyclicDungeonGenerator/layout_generation"
	"fmt"
)

type layoutVisualiser struct{}

func (l *layoutVisualiser) putCharArray(x, y int, c *[][]rune) {
	for i := 0; i < len(*c); i++ {
		for j := 0; j < len((*c)[0]); j++ {
			setcolorForRune((*c)[i][j])
			cw.PutChar((*c)[i][j], x+i, y+j)
		}
	}
}

func (l *layoutVisualiser) putMap(a *layout_generation.LayoutMap) {
	l.putCharArray(0, 0, a.WholeMapToCharArray())
}

func (l *layoutVisualiser) putInfo(a *layout_generation.LayoutMap, pattNum, desiredPNum int, fName, pName string, restarts int, rand bool) {
	sx, sy := a.GetSize()
	for y := 0; y < sy; y++ {
		for x := 0; x < sx; x++ {
			chr := a.GetCharOfElementAtCoords(x, y)
			setcolorForRune(chr)
			cw.PutChar(chr, x+sx*5+2, y)
		}
	}
	cw.SetColor(cw.BEIGE, cw.BLACK)
	cw.PutString(fmt.Sprintf("PATTERN SELECTED: #%d  ", desiredPNum), sx*5+2, sy+2)
	cw.PutString(fmt.Sprintf("PATTERN USED: #%d  ", pattNum), sx*5+2, sy+3)
	cw.PutString(fmt.Sprintf("FILE: %s  ", fName), sx*5+2, sy+4)
	cw.PutString(fmt.Sprintf("NAME: %s  ", pName), sx*5+2, sy+5)
	cw.PutString(fmt.Sprintf("%dx%d nodes", W, H), sx*5+2, sy+6)
	cw.PutString(fmt.Sprintf("Gen restarts: %d", restarts), sx*5+2, sy+7)
	if rand {
		cw.PutString("Random paths", sx*5+2, sy+8)
	} else {
		cw.PutString("Shortest paths", sx*5+2, sy+8)
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
