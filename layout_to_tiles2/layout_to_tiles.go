package layout_to_tiles2

import (
	"CyclicDungeonGenerator/layout_generation"
	"CyclicDungeonGenerator/random"
)

type LayoutToLevel struct {
	charmap      [][]rune
	roomW, roomH int
	rnd          *random.FibRandom
	CARoomChance, CAConnectionChance int
}

func (ltl *LayoutToLevel) Init(rnd *random.FibRandom, roomW, roomH int) {
	ltl.rnd = rnd
	ltl.roomW = roomW
	ltl.roomH = roomH
}

// roomSize is WITHOUT walls taken into account!
func (ltl *LayoutToLevel) MakeCharmap(layout *layout_generation.LayoutMap) *[][]rune {
	rw, rh := layout.GetSize()

	// +1 is for walls
	ltl.charmap = make([][]rune, rw*(ltl.roomW+1)+1)
	for i := range ltl.charmap {
		// +1 is for walls
		ltl.charmap[i] = make([]rune, rh*(ltl.roomH+1)+1)
	}

	for x := range ltl.charmap {
		for y := range ltl.charmap[x] {
			ltl.charmap[x][y] = ' ' // empty everything
		}
	}

	ltl.iterateNodes(layout, true, false)
	ltl.iterateNodes(layout, false, true)

	// draw perimeter walls
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

	ltl.iterateNodesForCA(layout)

	return &ltl.charmap
}

func (ltl *LayoutToLevel) iterateNodesForCA(layout *layout_generation.LayoutMap) {
	rw, rh := layout.GetSize()
	for lroomx := 0; lroomx < rw; lroomx++ {
		for lroomy := 0; lroomy < rh; lroomy++ {
			fromx := lroomx * (ltl.roomW+1)
			tox := (lroomx + 1) * (ltl.roomW+1)
			fromy := lroomy * (ltl.roomH+1)
			toy := (lroomy + 1) * (ltl.roomH+1)
			if layout.GetElement(lroomx, lroomy).IsNode() && ltl.rnd.RandomPercent() <= ltl.CARoomChance {
				ltl.dilateWalls(fromx, fromy, tox, toy, 1, 30)
				ltl.erodeWalls(fromx, fromy, tox, toy, 1, 30)
				ltl.dilateWalls(fromx, fromy, tox, toy, 1, 0)
			}
			if !layout.GetElement(lroomx, lroomy).IsNode() && ltl.rnd.RandomPercent() <= ltl.CAConnectionChance {
				ltl.dilateWalls(fromx, fromy, tox, toy, 1, 30)
				ltl.erodeWalls(fromx, fromy, tox, toy, 1, 30)
				ltl.dilateWalls(fromx, fromy, tox, toy, 1, 0)
			}
		}
	}
}

func (ltl *LayoutToLevel) iterateNodes(layout *layout_generation.LayoutMap, doConnections, doRooms bool) {
	rw, rh := layout.GetSize()
	for lroomx := 0; lroomx < rw; lroomx++ {
		for lroomy := 0; lroomy < rh; lroomy++ {
			layoutElem := layout.GetElement(lroomx, lroomy)
			// surround it with walls
			boundLeft := lroomx * (ltl.roomW + 1)
			boundRight := (lroomx + 1) * (ltl.roomW + 1)
			boundUpper := (lroomy) * (ltl.roomH + 1)
			boundLower := (lroomy + 1) * (ltl.roomH + 1)
			centerX := boundLeft + (ltl.roomW+1)/2
			centerY := boundUpper + (ltl.roomH+1)/2
			conns := layoutElem.GetAllConnectionsCoords()
			// is a room
			if doRooms && layoutElem.IsNode() {
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
					evenWAddition := 0
					evenHAddition := 0
					// for dealing with wrong doors placement for non-odd room sizes.
					if ltl.roomW % 2 == 0 {
						evenWAddition = 1
					}
					if ltl.roomH % 2 == 0 {
						evenHAddition = 1
					}
					if cx == 0 { // horiz
						centerXoff = centerX + ltl.rnd.RandInRange(-ltl.roomW/2 + evenWAddition, ltl.roomW/2)
					} else { // ver
						centerYoff = centerY + ltl.rnd.RandInRange(-ltl.roomH/2 + evenHAddition, ltl.roomH/2)
					}
					// for dealing with wrong doors placement for non-odd room sizes.
					if ltl.roomW % 2 == 0 && cx > 0 {
						centerXoff++
					}
					if ltl.roomH % 2 == 0 && cy > 0 {
						centerYoff++
					}
					connRune := '+'
					switch layoutElem.GetConnectionByCoords(cx, cy).LockNum {
					case 1:
						connRune = '%'
					case 2:
						connRune = '='
					}
					doorX := centerXoff+conns[connIndex][0]*(ltl.roomW+1)/2
					doorY := centerYoff+conns[connIndex][1]*(ltl.roomH+1)/2
					// restrict non-locked doors to be placed over locked doors:
					if connRune == '+' && ltl.charmap[doorX][doorY] != '#' {
						continue
					}
					ltl.charmap[doorX][doorY] = connRune
				}
			}
			// do nodes
			if doConnections && !layoutElem.IsNode() {
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
