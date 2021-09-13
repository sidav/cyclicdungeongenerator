package generator

type LayoutInterface interface {
	WholeMapToCharArray(bool) *[][]rune
	GetSize() (int, int)
	GetCharOfElementAtCoords(int, int) rune
}
