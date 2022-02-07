package generator

import "github.com/sidav/cyclicdungeongenerator/generator/layout_generation"

type LayoutInterface interface {
	WholeMapToCharArray(bool, bool, bool) *[][]rune
	GetSize() (int, int)
	GetElement(int, int) *layout_generation.Element
	GetCharOfElementAtCoords(int, int) rune
	CellToCharArray(int, int, bool, bool, bool) []rune
}
