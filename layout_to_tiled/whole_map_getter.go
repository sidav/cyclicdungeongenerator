package layout_to_tiled

import "CyclicDungeonGenerator/layout_generation"

type Tile struct {
	Char rune
}

func GetTileMap(a *layout_generation.LayoutMap) *[][]Tile {
	layoutw, layouth := a.GetSize()

	tilemap := make([][]Tile, layoutw*12)
	for i := range tilemap {
		tilemap[i] = make([]Tile, layouth*12)
	}

	for lx := 0; lx < layoutw; lx++ {
		for ly := 0; ly < layouth; ly++ {
			node := a.GetElement(lx, ly)
			conns := node.GetAllConnectionsCoords()
			if len(conns) > 0 {
				placeDoors := node.IsNode()
				roomStrs := getTilemapByNodeConnections(&conns, placeDoors)
				// roomSize := len(*roomStrs) - 1
				for rx := range(*roomStrs) {
					for ry := range((*roomStrs)[0]) {
						// print(rx, ry)
						tilemap[lx*12+rx][ly*12+ry] = Tile{Char: rune((*roomStrs)[ry][rx])}
					}
				}
			}
		}
	}
	return &tilemap
}
