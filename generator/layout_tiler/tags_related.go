package layout_tiler

import (
	"strconv"
	"strings"
)

func (ltl *LayoutTiler) getLayoutTagsForTileAtCoords(x, y int) []string {
	elem := ltl.layout.GetElement(x/(ltl.roomW+1), y/(ltl.roomH+1))
	if !elem.IsNode() {
		return []string{}
	}
	return elem.GetTags()
}

func (ltl *LayoutTiler) isTagUsedAtcoords(tag string, x, y int) bool {
	elem := ltl.layout.GetElement(x/(ltl.roomW+1), y/(ltl.roomH+1))
	if !elem.IsNode() {
		return tag == ""
	}
	tags := elem.GetTags()
	if len(tags) == 0 && tag == "" {
		return true 
	}
	for _, t := range tags {
		if t == tag {
			return true
		}
	}
	return false
}

func (ltl *LayoutTiler) getLockIdForTileAtCoords(x, y int) int {
	tags := ltl.getLayoutTagsForTileAtCoords(x, y)
	for _, tag := range tags {
		lockid, err := strconv.Atoi(strings.Replace(tag, "ky", "", -1))
		if err != nil {
			lockid, _ = strconv.Atoi(strings.Replace(tag, "key", "", -1))
		}
		return lockid
	}
	return 0
}

func (ltl *LayoutTiler) countTotalTagUsagesInLayout(tag string) int {
	usages := 0
	w, h := ltl.layout.GetSize()
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			elem := ltl.layout.GetElement(x, y)
			if elem.IsNode() {
				tagsOfElem := elem.GetTags()
				for _, t := range tagsOfElem {
					if strings.Contains(t, tag) {
						usages++
					}
				}
			}
		}
	}
	return usages
}

func (ltl *LayoutTiler) increaseTagUsages(tag string) {
	ltl.tagUsages[tag]++
}

func (ltl *LayoutTiler) getTagUsagesInTilemap(tag string) int {
	usages, found  := ltl.tagUsages[tag]
	if found {
		return usages
	}
	return 0
}

func (ltl *LayoutTiler) changeRandomTileWithTagTo(tag string, code tileCode, lockLevel int) {
	tile := ltl.getRandomFloorTileWithTag(tag)
	if tile != nil {
		tile.Code = code
		tile.LockId = lockLevel
	}
}

func (ltl *LayoutTiler) getRandomFloorTileWithTag(tag string) *Tile {
	tiles := make([]*Tile, 0)
	for x := range ltl.TileMap {
		for y := range ltl.TileMap[x] {
			if ltl.TileMap[x][y].Code == TILE_FLOOR && ltl.isTagUsedAtcoords(tag, x, y) {
				tiles = append(tiles, &ltl.TileMap[x][y])
			}
		}
	}
	if len(tiles) > 0 {
		return tiles[ltl.rnd.Rand(len(tiles))]
	}
	return nil
}

// TODO: rewrite all the following func
func (ltl *LayoutTiler) isSpaceEvenlyTagged(xx, yy, w, h int, tag string) bool {
	for x := xx; x < xx+w; x++ {
		for y := yy; y < yy+h; y++ {
			tileTag := ltl.getLayoutTagsForTileAtCoords(x, y)
			if len(tileTag) > 0 && tag == "" {
				return false
			}
			if !ltl.isTagUsedAtcoords(tag, x, y) { // strings.Contains(tileTag, tag) {
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
				ltl.TileMap[x][y].LockId = ltl.getLockIdForTileAtCoords(x, y)
			}
		}
	}
	// place entrypoint if ltl.TagForEntry is set
	if ltl.TagForEntry != "" && ltl.getTagUsagesInTilemap(ltl.TagForEntry) == 0 {
		ltl.changeRandomTileWithTagTo(ltl.TagForEntry, TILE_ENTRYPOINT, 0)
	}
	// place exitpoint if ltl.TagForEntry is set
	if ltl.TagForExit != "" && ltl.getTagUsagesInTilemap(ltl.TagForExit) == 0 {
		ltl.changeRandomTileWithTagTo(ltl.TagForExit, TILE_EXITPOINT, 0)
	}
	// place keys if ltl.TagsForKeys is non-empty
	if len(ltl.TagsForKeys) > 0 {
		for num, tag := range ltl.TagsForKeys {
			if ltl.getTagUsagesInTilemap(tag) == 0 {
				ltl.changeRandomTileWithTagTo(tag, TILE_KEY_PLACE, num+1)
			}
		}
	}
}
