package layout_tiler

import (
	"strconv"
	"strings"
)

func (ltl *LayoutTiler) getTagsForTileAtCoords(x, y int) string {
	elem := ltl.layout.GetElement(x/(ltl.roomW+1), y/(ltl.roomH+1))
	if !elem.IsNode() {
		return ""
	}
	return elem.GetTags()
}

func (ltl *LayoutTiler) countTotalTagUsagesInLayout(tag string) int {
	usages := 0
	w, h := ltl.layout.GetSize()
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			elem := ltl.layout.GetElement(x, y)
			if elem.IsNode() {
				if strings.Contains(elem.GetTags(), tag) {
					usages++
				}
			}
		}
	}
	return usages
}

// TODO: rewrite all the following func
func (ltl *LayoutTiler) isSpaceEvenlyTagged(xx, yy, w, h int, tag string) bool {
	for x := xx; x < xx+w; x++ {
		for y := yy; y < yy+h; y++ {
			tileTag := ltl.getTagsForTileAtCoords(x, y)
			if tileTag != "" && tag == "" {
				return false
			}
			if !strings.Contains(tileTag, tag) {
				return false
			}
		}
	}
	return true
}

func (ltl *LayoutTiler) finishTagsRelatedStuff() {
	// finalize keys: set lockId to key places
	for x := 0; x < len(ltl.TileMap); x++ {
		for y := 0; y < len(ltl.TileMap[0]); y++ {
			if ltl.TileMap[x][y].Code == TILE_KEY_PLACE {
				// TODO: consider multiple tags
				lockID, _ := strconv.Atoi(strings.Replace(ltl.getTagsForTileAtCoords(x, y), "ky", "", -1))
				ltl.TileMap[x][y].LockId = lockID
			}
		}
	}
}
