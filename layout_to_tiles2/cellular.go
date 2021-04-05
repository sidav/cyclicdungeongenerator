package layout_to_tiles2

func erodeWalls(iters, chancePerc int) {
	for i := 0; i < iters; i++ {
		var coordsToErode [][2]int
		for x := 1; x < len(charmap)-1; x++ {
			for y := 1; y < len(charmap[x])-1; y++ {
				if charmap[x][y] == '#' {
					adjDP, adjDC := countNeighbouring(x, y, '+')
					if adjDP+adjDC > 0 {
						continue
					}
					adjP, adjC := countNeighbouring(x, y, '#')
					adj := adjC + adjP
					if 5 <= adj && adj <= 7 && adjC < 3 {
						if rnd.RandomPercent() < chancePerc {
							coordsToErode = append(coordsToErode, [2]int{x, y})
							charmap[x][y] = ' '
						}
					}
				}
			}
		}
	}
	//for i := range coordsToErode {
	//		x, y := coordsToErode[i][0], coordsToErode[i][1]
	//		charmap[x][y] = '='
	//}
}

func dilateWalls(iters, chancePerc int) {
	var coords [][2]int
	for i := 0; i < iters; i++ {
		for x := 1; x < len(charmap)-1; x++ {
			for y := 1; y < len(charmap[x])-1; y++ {
				if charmap[x][y] == ' ' {
					adjDP, adjDC := countNeighbouring(x, y, '+')
					if adjDP+adjDC > 0 {
						continue
					}
					adjP, adjC := countNeighbouring(x, y, '#')
					adj := adjC + adjP
					if 3 <= adj && adjP > 0 {
						if rnd.RandomPercent() < chancePerc || adj >= 5 && adjP == 4 {
							coords = append(coords, [2]int{x, y})
							// charmap[x][y] = '#'
						}
					}
				}
			}
		}
	}
	for i := range coords {
			x, y := coords[i][0], coords[i][1]
			charmap[x][y] = '#'
	}
}


func countNeighbouring(xx, yy int, counts rune) (int, int) {
	countedPlus, countedCross := 0, 0
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			if x*y == 0 {
				if charmap[xx+x][yy+y] == counts {
					countedPlus++
				}
			}
			if x*y != 0 {
				if charmap[xx+x][yy+y] == counts {
					countedCross++
				}
			}
		}
	}
	return countedPlus, countedCross
}
