package layout_to_tiled_map

import (
	"CyclicDungeonGenerator/layout_generation"
	"CyclicDungeonGenerator/random"
)

type LayoutToLevel struct {
	TileMap                          [][]Tile
	submaps                          map[string][]submap
	roomW, roomH                     int
	rnd                              *random.FibRandom
	CARoomChance, CAConnectionChance int
	layout                           *layout_generation.LayoutMap
}

func (ltl *LayoutToLevel) Init(rnd *random.FibRandom, roomW, roomH int) {
	ltl.rnd = rnd
	ltl.roomW = roomW
	ltl.roomH = roomH
	ltl.submaps = make(map[string][]submap)
}

// roomSize is WITHOUT walls taken into account!
func (ltl *LayoutToLevel) ProcessLayout(layout *layout_generation.LayoutMap, submapsDir string) {
	ltl.layout = layout
	ltl.parseSubmapsDir(submapsDir)
	rw, rh := layout.GetSize()

	// +1 is for walls
	ltl.TileMap = make([][]Tile, rw*(ltl.roomW+1)+1)
	for i := range ltl.TileMap {
		// +1 is for walls
		ltl.TileMap[i] = make([]Tile, rh*(ltl.roomH+1)+1)
	}

	for x := range ltl.TileMap {
		for y := range ltl.TileMap[x] {
			ltl.TileMap[x][y].Code = TILE_FLOOR // empty everything
		}
	}

	ltl.iterateNodes(layout, true, false)
	ltl.iterateNodes(layout, false, true)

	// draw perimeter walls
	for x := range ltl.TileMap {
		for y := range ltl.TileMap[x] {
			//if ltl.TileMap[x][y].Code == TILE_NOT_SET {
			//	ltl.TileMap[x][y].Code = TILE_WALL
			//}
			if x == 0 || x == len(ltl.TileMap)-1 || y == 0 || y == len(ltl.TileMap[x])-1 {
				ltl.TileMap[x][y].Code = TILE_WALL
			}
		}
	}

	ltl.applySubmaps()
	ltl.iterateNodesForCA(layout)
	ltl.finishTagsRelatedStuff()
	// ltl.layout = nil // free memory
}

func (ltl *LayoutToLevel) GetCharMapForLevel() *[][]rune {
	w, h := len(ltl.TileMap), len(ltl.TileMap[0])
	rmap := make([][]rune, w)
	for i := range rmap {
		// +1 is for walls
		rmap[i] = make([]rune, h)
	}
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			rmap[x][y] = ltl.TileMap[x][y].GetChar()
		}
	}
	return &rmap
}

func (ltl *LayoutToLevel) iterateNodesForCA(layout *layout_generation.LayoutMap) {
	rw, rh := layout.GetSize()
	for lroomx := 0; lroomx < rw; lroomx++ {
		for lroomy := 0; lroomy < rh; lroomy++ {
			fromx := lroomx * (ltl.roomW + 1)
			tox := (lroomx + 1) * (ltl.roomW + 1)
			fromy := lroomy * (ltl.roomH + 1)
			toy := (lroomy + 1) * (ltl.roomH + 1)
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
							ltl.TileMap[x][y].Code = TILE_WALL
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
					if ltl.roomW%2 == 0 {
						evenWAddition = 1
					}
					if ltl.roomH%2 == 0 {
						evenHAddition = 1
					}
					if cx == 0 { // horiz
						centerXoff = centerX + ltl.rnd.RandInRange(-ltl.roomW/2+evenWAddition, ltl.roomW/2)
					} else { // ver
						centerYoff = centerY + ltl.rnd.RandInRange(-ltl.roomH/2+evenHAddition, ltl.roomH/2)
					}
					// for dealing with wrong doors placement for non-odd room sizes.
					if ltl.roomW%2 == 0 && cx > 0 {
						centerXoff++
					}
					if ltl.roomH%2 == 0 && cy > 0 {
						centerYoff++
					}

					doorX := centerXoff + conns[connIndex][0]*(ltl.roomW+1)/2
					doorY := centerYoff + conns[connIndex][1]*(ltl.roomH+1)/2

					currLockLevel := layoutElem.GetConnectionByCoords(cx, cy).LockNum
					// restrict non-locked doors to be placed over locked doors:
					if currLockLevel == 0 && ltl.TileMap[doorX][doorY].LockId != 0 {
						continue
					}
					ltl.TileMap[doorX][doorY].Code = TILE_DOOR
					ltl.TileMap[doorX][doorY].LockId = currLockLevel
				}
			}
			// do nodes
			if doConnections && !layoutElem.IsNode() {
				for x := boundLeft; x <= boundRight; x++ {
					for y := boundUpper; y <= boundLower; y++ {
						ltl.TileMap[x][y].Code = TILE_WALL
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
							ltl.TileMap[x][y].Code = TILE_FLOOR
						}
					}
				}
			}
		}
	}
}
