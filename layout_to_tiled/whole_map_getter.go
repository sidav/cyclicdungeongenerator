package layout_to_tiled

import "CyclicDungeonGenerator/layout_generation"

type Tile struct {
	Char rune
	placed bool
}

func twoVarsAreOfThatValues(c1, c2, e1, e2 rune) bool {
	if (c1 == e1 && c2 == e2) || (c1 == e2 && c2 == e1) {
		return true
	}
	return false
}

func selectTileForOverlap(c1, c2 rune) rune {
	if twoVarsAreOfThatValues(c1, c2, '+', '.') {
		return '+'
	}
	if twoVarsAreOfThatValues(c1, c2, '+', '#')  {
		return '+'
	}
	if twoVarsAreOfThatValues(c1, c2, '#', '.')  {
		return '#'
	}
	return c1
}

func TransformLayoutToTileMap(a *layout_generation.LayoutMap) *[][]Tile {
	layoutw, layouth := a.GetSize()

	tilemap := make([][]Tile, layoutw*11+1)
	for i := range tilemap {
		tilemap[i] = make([]Tile, layouth*11+1)
	}

	for lx := 0; lx < layoutw; lx++ {
		for ly := 0; ly < layouth; ly++ {
			node := a.GetElement(lx, ly)
			conns := node.GetAllConnectionsCoords()
			if len(conns) > 0 {
				placeDoors := node.IsNode()
				roomStrs := getTilemapByNodeConnections(&conns, placeDoors)
				// roomSize := len(*roomStrs) - 1
				for rx := range *roomStrs {
					for ry := range(*roomStrs)[0] {
						chr := rune((*roomStrs)[ry][rx])
						tilemap[lx*11+rx][ly*11+ry].Char = selectTileForOverlap(chr, tilemap[lx*11+rx][ly*11+ry].Char)
						tilemap[lx*11+rx][ly*11+ry].placed = true
					}
				}
			}
		}
	}
	return &tilemap
}
