package layout_generation

func (l *LayoutMap) PerformLocksCheckForPattern(p *pattern) bool {
	// check locks
	for _, i := range p.instructions {
		if i.lockNumber > 0 {
			w, h := l.GetSize()
			lockFound := false
			iterateCells:
			for x := 0; x < w; x++ {
				for y := 0; y < h; y++ {
					for _, conn := range l.elements[x][y].connections {
						if conn != nil && conn.LockNum == i.lockNumber {
							lockFound = true
							break iterateCells
						}
					}
				}
			}
			if !lockFound {
				return false
			}
		}
	}
	return true
}
