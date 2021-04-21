package layout_to_tiled_map

import "strconv"

type tileCode uint8

const (
	TILE_FLOOR tileCode = iota
	TILE_WALL
	TILE_DOOR
	TILE_WATER
	TILE_ENTRYPOINT
	TILE_EXITPOINT
	TILE_NOT_SET
)

var CharToTileCode = map[rune]tileCode{
	' ': TILE_FLOOR,
	'#': TILE_WALL,
	'+': TILE_DOOR,
	'~': TILE_WATER,
	'<': TILE_ENTRYPOINT,
	'>': TILE_EXITPOINT,
}

type Tile struct {
	Code   tileCode
	LockId int
}

func (t *Tile) GetChar() rune {
	if t.Code == TILE_DOOR && t.LockId != 0 {
		return rune(strconv.Itoa(t.LockId)[0])
	}
	for i := range CharToTileCode {
		if CharToTileCode[i] == t.Code {
			return i
		}
	}
	return '?'
}
