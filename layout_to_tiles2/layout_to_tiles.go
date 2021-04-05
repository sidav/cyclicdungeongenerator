package layout_to_tiles2

import (
	"CyclicDungeonGenerator/layout_generation"
	"CyclicDungeonGenerator/random"
)

type LayoutToLevel struct {
	charmap  [][]rune
	roomSize int
	rnd *random.FibRandom
}

func (ltl *LayoutToLevel) Init(rnd *random.FibRandom, roomSize int) {
	ltl.rnd = rnd
	ltl.roomSize = roomSize
}

// roomSize is WITHOUT walls taken into account!
func (ltl *LayoutToLevel) MakeCharmap(layout *layout_generation.LayoutMap) [][]rune {
	rw, rh := layout.GetSize()

	// +1 is for walls
	ltl.charmap = make([][]rune, rw*(ltl.roomSize+1)+1)
	for i := range ltl.charmap {
		// +1 is for walls
		ltl.charmap[i] = make([]rune, rh*(ltl.roomSize+1)+1)
	}

	for x := range ltl.charmap {
		for y := range ltl.charmap[x] {
			ltl.charmap[x][y] = ' ' // empty everything
		}
	}

	ltl.doForConnections(layout)
	ltl.doForRooms(layout)

	// draw perimeter walls, replace "temp walls" with true walls
	for x := range ltl.charmap {
		for y := range ltl.charmap[x] {
			if ltl.charmap[x][y] == '?' {
				ltl.charmap[x][y] = '#'
			}
			if x == 0 || x == len(ltl.charmap)-1 || y == 0 || y == len(ltl.charmap[x])-1 {
				ltl.charmap[x][y] = '#'
			}
		}
	}
	ltl.dilateWalls(2, 30)
	ltl.erodeWalls(1, 10)
	ltl.dilateWalls(1, 0)

	return ltl.charmap
}

func (ltl *LayoutToLevel) doForRooms(layout *layout_generation.LayoutMap) {
	rw, rh := layout.GetSize()
	for lroomx := 0; lroomx < rw; lroomx++ {
		for lroomy := 0; lroomy < rh; lroomy++ {
			layoutElem := layout.GetElement(lroomx, lroomy)
			// surround it with walls
			boundLeft := lroomx * (ltl.roomSize + 1)
			boundRight := (lroomx + 1) * (ltl.roomSize + 1)
			boundUpper := (lroomy) * (ltl.roomSize + 1)
			boundLower := (lroomy + 1) * (ltl.roomSize + 1)
			centerX := boundLeft + (ltl.roomSize+1)/2
			centerY := boundUpper + (ltl.roomSize+1)/2
			conns := layoutElem.GetAllConnectionsCoords()
			// is a room
			if layoutElem.IsNode() {
				for x := boundLeft; x <= boundRight; x++ {
					for y := boundUpper; y <= boundLower; y++ {
						if x == boundRight || x == boundLeft || y == boundUpper || y == boundLower {
							ltl.charmap[x][y] = '#'
						}
					}
				}
				for connIndex := range conns {
					cx, cy := conns[connIndex][0], conns[connIndex][1]
					// randomly displace center for creating random door offset
					centerXoff := centerX
					centerYoff := centerY
					if cx == 0 { // horiz
						centerXoff = centerX + ltl.rnd.RandInRange(-ltl.roomSize/2, ltl.roomSize/2)
					} else { // ver
						centerYoff = centerY + ltl.rnd.RandInRange(-ltl.roomSize/2, ltl.roomSize/2)
					}
					connRune := '+'
					switch layoutElem.GetConnectionByCoords(cx, cy).LockNum {
					case 1:
						connRune = '%'
					case 2:
						connRune = '='
					}
					ltl.charmap[centerXoff+conns[connIndex][0]*(ltl.roomSize+1)/2][centerYoff+conns[connIndex][1]*(ltl.roomSize+1)/2] = connRune
				}
			}
		}
	}
}

func (ltl *LayoutToLevel) doForConnections(layout *layout_generation.LayoutMap) {
	rw, rh := layout.GetSize()
	for lroomx := 0; lroomx < rw; lroomx++ {
		for lroomy := 0; lroomy < rh; lroomy++ {
			layoutElem := layout.GetElement(lroomx, lroomy)
			// surround it with walls
			boundLeft := lroomx * (ltl.roomSize + 1)
			boundRight := (lroomx + 1) * (ltl.roomSize + 1)
			boundUpper := (lroomy) * (ltl.roomSize + 1)
			boundLower := (lroomy + 1) * (ltl.roomSize + 1)
			conns := layoutElem.GetAllConnectionsCoords()
			// is a room
			if !layoutElem.IsNode() {
				for x := boundLeft; x <= boundRight; x++ {
					for y := boundUpper; y <= boundLower; y++ {
						ltl.charmap[x][y] = '#'
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
					for x := boundLeft + 1 - leftOff; x < boundRight+rightOff; x++ {
						for y := boundUpper + 1 - upOff; y < boundLower+botOff; y++ {
							ltl.charmap[x][y] = ' '
						}
					}
				}
			}
		}
	}
}
