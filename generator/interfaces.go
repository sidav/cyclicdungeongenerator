package generator

type LayoutInterface interface {
	WholeMapToCharArray() *[][]rune
	GetSize() (int, int)
	GetCharOfElementAtCoords(int, int) rune
}
