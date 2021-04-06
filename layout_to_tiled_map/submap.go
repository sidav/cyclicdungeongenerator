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
	for tries := 0; tries < 3; tries++ {
		indexOffset := ltl.rnd.Rand(len(ltl.submaps))
		for i := range ltl.submaps {
			ind := (i + indexOffset) % len(ltl.submaps)
			if ltl.submaps[ind].timesUsed == 0 {
				ltl.applySubmapAtRandom(&ltl.submaps[ind])
			}
		}
	}
}

func (ltl *LayoutToLevel) applySubmapAtRandom(sm *submap) {
	sm.rotate(ltl.rnd.Rand(4))
	smH, smW := len(sm.chars), len(sm.chars[0])
	applicableCoords := make([][2]int, 0)
	for x := 0; x < len(ltl.charmap)-smW; x++ {
		for y := 0; y < len(ltl.charmap[x])-smH; y++ {
			if ltl.isSpaceEmpty(x, y, smW, smH) {
				applicableCoords = append(applicableCoords, [2]int{x, y})
			}
		}
	}
	if len(applicableCoords) > 0 {
		randCoordsIndex := ltl.rnd.Rand(len(applicableCoords))
		ltl.applySubmapAtCoords(sm, applicableCoords[randCoordsIndex][0], applicableCoords[randCoordsIndex][1])
		sm.timesUsed++
	}
}


func (ltl *LayoutToLevel) applySubmapAtCoords(sm *submap, xx, yy int) bool {
	smH, smW := len(sm.chars), len(sm.chars[0])
	for x := 0; x < smW; x++ {
		for y := 0; y < smH; y++ {
			ltl.charmap[xx+x][yy+y] = sm.chars[y][x]
		}
	}
	return true
}

func (ltl *LayoutToLevel) isSpaceEmpty(xx, yy, w, h int) bool {
	for x := xx; x < xx+w; x++ {
		for y := yy; y < yy+h; y++ {
			if ltl.charmap[x][y] != ' ' {
				return false
			}
		}
	}
	return true
}
