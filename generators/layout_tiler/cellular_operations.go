package layout_tiler

func (ltl *LayoutTiler) erodeWalls(fromx, fromy, tox, toy, iters, chancePerc int) {
	var coordsToErode [][2]int
	for i := 0; i < iters; i++ {
		for x := fromx; x < tox; x++ {
			for y := fromy; y < toy; y++ {
				if ltl.TileMap[x][y].Code == TILE_WALL {
					adjDP, adjDC := ltl.countDoorsNearby(x, y)
					if adjDP+adjDC > 0 {
						continue
					}
					adjP, adjC := ltl.countNeighbouring(x, y, TILE_WALL)
					adj := adjC + adjP
					if 5 <= adj && adj <= 7 && adjP < 3 {
						if ltl.rnd.RandomPercent() < chancePerc {
							coordsToErode = append(coordsToErode, [2]int{x, y})
							// ltl.TileMap[x][y] = ' '
						}
					}
				}
			}
		}
	}
	for i := range coordsToErode {
		x, y := coordsToErode[i][0], coordsToErode[i][1]
		ltl.TileMap[x][y].Code = TILE_FLOOR
	}
}

func (ltl *LayoutTiler) dilateWalls(fromx, fromy, tox, toy, iters, chancePerc int) {
	var coords [][2]int
	for i := 0; i < iters; i++ {
		for x := fromx; x < tox; x++ {
			for y := fromy; y < toy; y++ {
				if ltl.TileMap[x][y].Code == TILE_FLOOR {
					adjDP, _ := ltl.countDoorsNearby(x, y)
					if adjDP > 0 {
						continue
					}
					adjP, adjC := ltl.countNeighbouring(x, y, TILE_WALL)
					adj := adjC + adjP
					if 3 <= adj && adjP > 0 {
						if ltl.rnd.RandomPercent() < chancePerc || adj >= 5 && adjP == 4 {
							coords = append(coords, [2]int{x, y})
							// ltl.TileMap[x][y] = '#'
						}
					}
				}
			}
		}
	}
	for i := range coords {
		x, y := coords[i][0], coords[i][1]
		ltl.TileMap[x][y].Code = TILE_WALL
	}
}

func (ltl *LayoutTiler) countDoorsNearby(xx, yy int) (int, int) {
	countedPlus, countedCross := 0, 0
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			if xx+x < 0 || xx+x >= len(ltl.TileMap) || yy+y < 0 || yy+y >= len(ltl.TileMap[0]) {
				continue
			}
			if ltl.TileMap[xx+x][yy+y].Code == TILE_DOOR {
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

func (ltl *LayoutTiler) countNeighbouring(xx, yy int, counts tileCode) (int, int) {
	countedPlus, countedCross := 0, 0
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			if xx+x < 0 || xx+x >= len(ltl.TileMap) || yy+y < 0 || yy+y >= len(ltl.TileMap[0]) {
				continue
			}
			if counts == ltl.TileMap[xx+x][yy+y].Code {
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
