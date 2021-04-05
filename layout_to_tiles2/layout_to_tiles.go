package layout_to_tiles2

import (
	"CyclicDungeonGenerator/layout_generation"
	"CyclicDungeonGenerator/random"
)

var charmap [][]rune
var roomsize int
var rnd *random.FibRandom

// roomsize is WITHOUT walls taken into account!
func MakeCharmap(rndgen *random.FibRandom, roomSize int, layout *layout_generation.LayoutMap) [][]rune {
	roomsize = roomSize
	rw, rh := layout.GetSize()
	rnd = rndgen

	// +1 is for walls
	charmap = make([][]rune, rw*(roomsize+1)+1)
	for i := range charmap {
		// +1 is for walls
		charmap[i] = make([]rune, rh*(roomsize+1)+1)
	}

	for x := range charmap {
		for y := range charmap[x] {
			charmap[x][y] = ' ' // empty everything
		}
	}

	doForConnections(layout)
	doForRooms(layout)


	// draw perimeter walls, replace "temp walls" with true walls
	for x := range charmap {
		for y := range charmap[x] {
			if charmap[x][y] == '?' {
				charmap[x][y] = '#'
			}
			if x == 0 || x == len(charmap)-1 || y == 0 || y == len(charmap[x])-1 {
				charmap[x][y] = '#'
			}
		}
	}
	dilateWalls(2,30)
	erodeWalls(1, 10)
	dilateWalls(1,0)

	return charmap
}

func doForRooms(layout *layout_generation.LayoutMap) {
	rw, rh := layout.GetSize()
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
					cx, cy := conns[connIndex][0], conns[connIndex][1]
					// randomly displace center for creating random door offset
					centerXoff := centerX
					centerYoff := centerY
					if cx == 0 {  // horiz
						centerXoff = centerX + rnd.RandInRange(-roomsize/2, roomsize/2)
					} else { // ver
						centerYoff = centerY + rnd.RandInRange(-roomsize/2, roomsize/2)
					}
					connRune := '+'
					switch layoutElem.GetConnectionByCoords(cx, cy).LockNum {
					case 1: connRune = '%'
					case 2: connRune = '='
					}
					charmap[centerXoff+conns[connIndex][0]*(roomsize+1)/2][centerYoff+conns[connIndex][1]*(roomsize+1)/2] = connRune
				}
			}
		}
	}
}

func doForConnections(layout *layout_generation.LayoutMap) {
	rw, rh := layout.GetSize()
	for lroomx := 0; lroomx < rw; lroomx++ {
		for lroomy := 0; lroomy < rh; lroomy++ {
			layoutElem := layout.GetElement(lroomx, lroomy)
			// surround it with walls
			boundLeft := lroomx * (roomsize + 1)
			boundRight := (lroomx + 1) * (roomsize + 1)
			boundUpper := (lroomy) * (roomsize + 1)
			boundLower := (lroomy + 1) * (roomsize + 1)
			conns := layoutElem.GetAllConnectionsCoords()
			// is a room
			if !layoutElem.IsNode() {
				for x := boundLeft; x <= boundRight; x++ {
					for y := boundUpper; y <= boundLower; y++ {
						charmap[x][y] = '#'
					}
				}
				for connIndex := range conns {
					conx, cony := conns[connIndex][0], conns[connIndex][1]
					leftOff := 0
					rightOff := 0
					upOff := 0
					botOff := 0
					if conx == -1 {
						leftOff = 1
					}
					if conx == 1 {
						rightOff = 1
					}
					if cony == -1 {
						upOff = 1
					}
					if cony == 1 {
						botOff = 1
					}
					for x := boundLeft+1-leftOff; x < boundRight+rightOff; x++ {
						for y := boundUpper+1-upOff; y < boundLower+botOff; y++ {
							charmap[x][y] = ' '
						}
					}
				}
			}
		}
	}
}
