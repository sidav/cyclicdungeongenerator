package main

import (
	"CyclicDungeonGenerator/cyc_dung_gen"
	cw "TCellConsoleWrapper"
)

func main() {
	cw.Init_console()
	defer cw.Close_console()

	generatedMap := cyc_dung_gen.Generate()
	putMap(generatedMap)
	cw.ReadKey()
}

func putMap(a *[][]rune) {
	for y:=0;y<len(*a);y++ {
		for x:=0;x<len((*a)[0]);x++{
			chr := (*a)[x][y]
			switch chr {
			case '.':
				cw.SetFgColor(cw.DARK_GRAY)
			case '*':
				cw.SetFgColor(cw.DARK_RED)
			default:
				cw.SetFgColor(cw.CYAN)
			}
			cw.PutChar(chr, x, y)
		}
	}
}
