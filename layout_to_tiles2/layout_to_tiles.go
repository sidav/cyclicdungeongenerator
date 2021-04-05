package layout_to_tiles2

import "CyclicDungeonGenerator/layout_generation"

// roomsize is WITHOUT walls taken into account!
func MakeCharmap(roomsize int, layout *layout_generation.LayoutMap) [][]rune {
	rw, rh := layout.GetSize()

	// +1 is for walls
	charmap := make([][]rune, rw*(roomsize+1)+1)
	for i := range charmap {
		// +1 is for walls
		charmap[i] = make([]rune, rh*(roomsize+1)+1)
	}

	for x := range charmap {
		for y := range charmap[x] {
			charmap[x][y] = ' ' // empty everything
		}
	}

	for lroomx := 0; lroomx < rw; lroomx++ {
		for lroomy := 0; lroomy < rh; lroomy++ {
			layoutElem := layout.GetElement(lroomx, lroomy)
			// surround it with walls
			boundLeft := lroomx * (roomsize + 1)
			boundRight := (lroomx + 1) * (roomsize + 1)
			boundUpper := (lroomy) * (roomsize + 1)
			boundLower := (lroomy + 1) * (roomsize + 1)
			centerX := boundLeft + (roomsize+1)/2
			centerY := boundUpper + (roomsize+1)/2
			conns := layoutElem.GetAllConnectionsCoords()
			// is a room
			if layoutElem.IsNode() {
				for x := boundLeft; x <= boundRight; x++ {
					for y := boundUpper; y <= boundLower; y++ {
						if x == boundRight || x == boundLeft || y == boundUpper || y == boundLower {
							charmap[x][y] = '#'
						}
					}
				}
				for connIndex := range conns {
					charmap[centerX+conns[connIndex][0]*(roomsize+1)/2][centerY + conns[connIndex][1]*(roomsize+1)/2] = '+'
				}
			} else {
				if len(conns) == 0 {
					for x := boundLeft; x <= boundRight; x++ {
						for y := boundUpper; y <= boundLower; y++ {
							charmap[x][y] = '#'
						}
					}
				} else {
					for connIndex := range conns {
						for x := boundLeft; x <= boundRight; x++  {
							for y := boundUpper; y <= boundLower; y++ {
								if x == centerX+conns[connIndex][0]*(roomsize+1)/2 || y == centerY + conns[connIndex][1]*(roomsize+1)/2 {
									// charmap[x][y] = '#'
								}
							}
						}
					}
				}
			}
		}
	}

	// draw perimeter walls
	for x := range charmap {
		for y := range charmap[x] {
			if x == 0 || x == len(charmap)-1 || y == 0 || y == len(charmap[x])-1 {
				charmap[x][y] = '#'
			}
		}
	}
	return charmap
}
