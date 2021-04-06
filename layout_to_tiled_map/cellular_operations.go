package layout_to_tiled_map

func (ltl *LayoutToLevel) erodeWalls(fromx, fromy, tox, toy, iters, chancePerc int) {
	var coordsToErode [][2]int
	for i := 0; i < iters; i++ {
		for x := fromx; x < tox; x++ {
			for y := fromy; y < toy; y++ {
				if ltl.charmap[x][y] == '#' {
					adjDP, adjDC := ltl.countDoorsNearby(x, y)
					if adjDP+adjDC > 0 {
						continue
					}
					adjP, adjC := ltl.countNeighbouring(x, y, '#')
					adj := adjC + adjP
					if 5 <= adj && adj <= 7 && adjP < 3 {
						if ltl.rnd.RandomPercent() < chancePerc {
							coordsToErode = append(coordsToErode, [2]int{x, y})
							// ltl.charmap[x][y] = ' '
						}
					}
				}
			}
		}
	}
	for i := range coordsToErode {
		x, y := coordsToErode[i][0], coordsToErode[i][1]
		ltl.charmap[x][y] = ' '
	}
}

func (ltl *LayoutToLevel) dilateWalls(fromx, fromy, tox, toy, iters, chancePerc int) {
	var coords [][2]int
	for i := 0; i < iters; i++ {
		for x := fromx; x < tox; x++ {
			for y := fromy; y < toy; y++ {
				if ltl.charmap[x][y] == ' ' {
					adjDP, _ := ltl.countDoorsNearby(x, y)
					if adjDP > 0 {
						continue
					}
					adjP, adjC := ltl.countNeighbouring(x, y, '#')
					adj := adjC + adjP
					if 3 <= adj && adjP > 0 {
						if ltl.rnd.RandomPercent() < chancePerc || adj >= 5 && adjP == 4 {
							coords = append(coords, [2]int{x, y})
							// ltl.charmap[x][y] = '#'
						}
					}
				}
			}
		}
	}
	for i := range coords {
		x, y := coords[i][0], coords[i][1]
		ltl.charmap[x][y] = '#'
	}
}

func (ltl *LayoutToLevel) countDoorsNearby(xx, yy int) (int, int) {
	countedPlus, countedCross := 0, 0
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			if xx+x < 0 || xx+x >= len(ltl.charmap) || yy+y < 0 || yy+y >= len(ltl.charmap[0]) {
				continue
			}
			if '+' == ltl.charmap[xx+x][yy+y] || '%' == ltl.charmap[xx+x][yy+y] || '=' == ltl.charmap[xx+x][yy+y] {
				if x*y == 0 {
					countedPlus++
				}
				if x*y != 0 {
					countedCross++
				}
			}
		}
	}
	return countedPlus, countedCross
}

func (ltl *LayoutToLevel) countNeighbouring(xx, yy int, counts rune) (int, int) {
	countedPlus, countedCross := 0, 0
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			if xx+x < 0 || xx+x >= len(ltl.charmap) || yy+y < 0 || yy+y >= len(ltl.charmap[0]) {
				continue
			}
			if counts == ltl.charmap[xx+x][yy+y] {
				if x*y == 0 {
					countedPlus++
				}
				if x*y != 0 {
					countedCross++
				}
			}
		}
	}
	return countedPlus, countedCross
}
