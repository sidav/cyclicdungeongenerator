package layout_generation

import "strconv"

// output. Used in benchmark to detect uniquity, should not be removed.

func (lm *LayoutMap) GetCharOfElementAtCoords(x, y int) rune { // just for rendering
	elem := lm.elements[x][y]
	// rune := '?'
	if elem.isEmpty() {
		return '.'
	}
	if elem.isObstacle {
		return '#'
	}
	if elem.IsNode() {
		if elem.GetName() == "" {
			return 'R'
		}
		return rune(elem.nodeInfo.nodeName[0])
	}
	if elem.isPartOfAPath() {
		number := elem.pathInfo.pathNumber
		return rune(strconv.Itoa(number)[0])
	}
	return '?'
}

func (lm *LayoutMap) CellToCharArray(cellx, celly int, renderPathNumbers bool) [][]rune {
	e := lm.elements[cellx][celly]
	ca := make([][]rune, 5)
	for i := range ca {
		ca[i] = make([]rune, 5)
	}

	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			ca[x][y] = '#'
		}
	}
	// draw node
	if e.nodeInfo != nil {
		for x := 1; x < 4; x++ {
			for y := 1; y < 4; y++ {
				ca[x][y] = ' '
			}
		}
		for x := -1; x <= 1; x++ {
			for y := -1; y <= 1; y++ {
				conn := e.GetConnectionByCoords(x, y)
				if conn != nil {
					if conn.IsLocked {
						if conn.LockNum == 1 {
							ca[2+x*2][2+y*2] = '%'
						} else {
							ca[2+x*2][2+y*2] = '='
						}
					} else {
						if conn.IsNodeExtension {
							ca[2+x*2-y][2+y*2] = ' '
							ca[2+x*2][2+y*2] = ' '
							ca[2+x*2+y][2+y*2] = ' '
							ca[2+x*2][2+y*2-x] = ' '
							ca[2+x*2][2+y*2+x] = ' '
						} else {
							ca[2+x*2][2+y*2] = '+'
						}
					}
				}
			}
		}
		if e.GetName() != "" {
			ca[1][2] = rune(e.nodeInfo.nodeName[0])
			ca[2][2] = rune(e.nodeInfo.nodeName[1])
			ca[3][2] = rune(e.nodeInfo.nodeName[2])
		}
		if renderPathNumbers && e.pathInfo != nil {
			ca[2][1] = rune(strconv.Itoa(e.pathInfo.pathNumber)[0])
		}
		switch len(e.nodeInfo.nodeTags) {
		case 1:
			if len(e.nodeInfo.nodeTags[0]) > 2 {
				ca[1][3] = rune(e.nodeInfo.nodeTags[0][0])
				ca[2][3] = rune(e.nodeInfo.nodeTags[0][1])
				ca[3][3] = rune(e.nodeInfo.nodeTags[0][2])
			}
		case 2:
			ca[1][3] = rune(e.nodeInfo.nodeTags[0][0])
			ca[2][3] = '-'
			ca[3][3] = rune(e.nodeInfo.nodeTags[1][0])
		case 3:
			ca[1][3] = rune(e.nodeInfo.nodeTags[0][0])
			ca[2][3] = rune(e.nodeInfo.nodeTags[1][0])
			ca[3][3] = rune(e.nodeInfo.nodeTags[2][0])
		}
		// draw path cell
	} else if e.pathInfo != nil {
		if renderPathNumbers {
			ca[2][2] = rune(strconv.Itoa(e.pathInfo.pathNumber)[0])
		} else {
			ca[2][2] = ' '
		}
		for x := -1; x <= 1; x++ {
			for y := -1; y <= 1; y++ {
				if e.GetConnectionByCoords(x, y) != nil {
					ca[2+x*2][2+y*2] = ' '
					ca[2+x][2+y] = ' '
				}
			}
		}
	}
	return ca
}

func (lm *LayoutMap) WholeMapToCharArray(pathNumbers bool) *[][]rune {
	sx, sy := lm.GetSize()
	ca := make([][]rune, 5*sx)
	for i := range ca {
		ca[i] = make([]rune, 5*sy)
	}
	for x := 0; x < len(lm.elements); x++ {
		for y := 0; y < len(lm.elements[0]); y++ {
			cellArr := lm.CellToCharArray(x, y, pathNumbers)
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					ca[5*x+i][5*y+j] = cellArr[i][j]
				}
			}
		}
	}
	return &ca
}
