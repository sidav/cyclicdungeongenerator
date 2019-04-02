package main

import (
	"CyclicDungeonGenerator/layout_generation"
	cw "TCellConsoleWrapper"
	rnd "github.com/sidav/golibrl/random"
)

func main() {
	rnd.Randomize()
	//for i:=0;i<10000;i++ {
	//	x, y := rnd.RandomCoordsInRangeFrom(40, 12, 10)
	//	cw.PutChar('#', x, y)
	//}
	//cw.Flush_console()
	//cw.ReadKey()

	generatedMap := layout_generation.Generate()

	cw.Init_console()
	defer cw.Close_console()

	putMap(generatedMap)
	cw.Flush_console()
	cw.ReadKey()
}

func putCharArray(x, y int, c [][]rune) {
	for i := 0; i < len(c); i++ {
		for j := 0; j < len(c[0]); j++ {
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


//func putMap(a *layout_generation.LayoutMap) {
//	sx, sy := a.GetSize()
//	for y := 0; y < sy; y++ {
//		for x := 0; x < sx; x++ {
//			chr := a.GetCharOfElementAtCoords(x, y)
//			switch chr {
//			case 'S', 'F', 'N':
//				cw.SetFgColor(cw.GREEN)
//			case '1', '2', '3', '4':
//				cw.SetFgColor(cw.DARK_CYAN)
//			case '.':
//				cw.SetFgColor(cw.BEIGE)
//			case '#':
//				cw.SetFgColor(cw.DARK_RED)
//			default:
//				cw.SetFgColor(cw.BLUE)
//			}
//			cw.PutChar(chr, x, y)
//		}
//	}
//}
