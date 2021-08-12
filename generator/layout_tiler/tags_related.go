package layout_tiler

import (
	"strconv"
	"strings"
)

func (ltl *LayoutTiler) getLayoutTagsForTileAtCoords(x, y int) string {
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
			if ltl.TileMap[x][y].Code == TILE_FLOOR && ltl.getLayoutTagsForTileAtCoords(x, y) == tag {
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
				lockID, _ := strconv.Atoi(strings.Replace(ltl.getLayoutTagsForTileAtCoords(x, y), "ky", "", -1))
				ltl.TileMap[x][y].LockId = lockID
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
