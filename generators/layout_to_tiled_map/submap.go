package layout_to_tiled_map

type submap struct {
	chars     [][]rune
	timesUsed int
}

func (sm *submap) rotate(times int) {
	for t := 0; t < times; t++ {
		var newChars [][]rune
		for y := 0; y < len(sm.chars[0]); y++ {
			var arr []rune
			for x := 0; x < len(sm.chars); x++ {
				arr = append(arr, sm.chars[x][y])
			}
			newChars = append(newChars, arr)
		}
		sm.chars = newChars
	}
}

func (ltl *LayoutToLevel) applySubmaps() {
	const TRIES_FOR_SUBMAP_PLACEMENT = 3
	// iterate through tags
	for tag := range ltl.submaps {
		totalSubmapUsesForTag := 0
		maxTagUses := ltl.countTotalTagUsagesInLayout(tag)

		for tries := 0; tries < TRIES_FOR_SUBMAP_PLACEMENT; tries++ {
			indexOffset := ltl.rnd.Rand(len(ltl.submaps[tag]))
			for i := range ltl.submaps[tag] {
				placed := false
				if tag != "" && totalSubmapUsesForTag == maxTagUses {
					break
				}
				ind := (i + indexOffset) % len(ltl.submaps[tag])
				if ltl.submaps[tag][ind].timesUsed == 0 {
					placed = ltl.applySubmapAtRandom(&ltl.submaps[tag][ind], tag)
				}
				if placed {
					totalSubmapUsesForTag++
				}
			}
		}
	}
}

func (ltl *LayoutToLevel) applySubmapAtRandom(sm *submap, tag string) bool {
	sm.rotate(ltl.rnd.Rand(4))
	smH, smW := len(sm.chars), len(sm.chars[0])
	applicableCoords := make([][2]int, 0)
	for x := 0; x < len(ltl.TileMap)-smW; x++ {
		for y := 0; y < len(ltl.TileMap[x])-smH; y++ {
			if ltl.isSpaceEmpty(x, y, smW, smH) && ltl.isSpaceEvenlyTagged(x, y, smW, smH, tag) {
				applicableCoords = append(applicableCoords, [2]int{x, y})
			}
		}
	}
	if len(applicableCoords) > 0 {
		randCoordsIndex := ltl.rnd.Rand(len(applicableCoords))
		ltl.applySubmapAtCoords(sm, applicableCoords[randCoordsIndex][0], applicableCoords[randCoordsIndex][1])
		sm.timesUsed++
		return true
	}
	return false
}

func (ltl *LayoutToLevel) applySubmapAtCoords(sm *submap, xx, yy int) {
	smH, smW := len(sm.chars), len(sm.chars[0])
	for x := 0; x < smW; x++ {
		for y := 0; y < smH; y++ {
			code, set := CharToTileCode[sm.chars[y][x]]
			if set {
				ltl.TileMap[xx+x][yy+y].Code = code
			} else {
				ltl.TileMap[xx+x][yy+y].Code = TILE_NOT_SET
			}
		}
	}
}

func (ltl *LayoutToLevel) isSpaceEmpty(xx, yy, w, h int) bool {
	for x := xx; x < xx+w; x++ {
		for y := yy; y < yy+h; y++ {
			if ltl.TileMap[x][y].Code != TILE_FLOOR {
				return false
			}
		}
	}
	return true
}
