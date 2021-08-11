package main

import (
	cw "cyclicdungeongenerator/console_wrapper"
	layout_generation2 "cyclicdungeongenerator/generators/layout_generation"
	"cyclicdungeongenerator/generators/layout_tiler"
	"cyclicdungeongenerator/random"
	"fmt"
	"strconv"
)

type tiledMapVisualiser struct {
	roomW, roomH                int
	drawRoomNames, drawRoomTags bool
}

func (g *tiledMapVisualiser) convertLayoutToLevelAndDraw(rnd *random.FibRandom, layout *layout_generation2.LayoutMap) {
	cw.Clear_console()
	tileMap := genWrapper.ConvertLayoutToTiledMap(rnd, layout, g.roomW, g.roomH, "generators/layout_tiler/submaps")
	g.drawLevel(&tileMap, 0, 0)
	rw, rh := layout.GetSize()

	if g.drawRoomTags || g.drawRoomNames {
		for rx := 0; rx < rw; rx++ {
			for ry := 0; ry < rh; ry++ {
				node := layout.GetElement(rx, ry)
				conns := node.GetAllConnectionsCoords()
				if len(conns) > 0 {
					cw.SetFgColor(cw.GREEN)
					if node.IsNode() {
						if g.drawRoomNames {
							name := node.GetName()
							strlen := len(name)
							offset := (g.roomW+1)/2 - strlen/2
							cw.PutString(name, rx*(g.roomW+1)+offset, ry*(g.roomH+1)+(g.roomH+1)/2)
						}
						if g.drawRoomTags {
							tags := node.GetTags()
							strlen := len(tags)
							offset := (g.roomW+1)/2 - strlen/2
							cw.PutString(tags, rx*(g.roomW+1)+offset, ry*(g.roomH+1)+(g.roomH+1)/2+1)
						}
					}
				}
			}
		}
	}
}

func (g *tiledMapVisualiser) putInfo(a *layout_generation2.LayoutMap, pattNum, desiredPNum int, fName, pName string, restarts, maxDesiredRestarts int, rand bool) {
	sx, sy := a.GetSize()
	for y := 0; y < sy; y++ {
		for x := 0; x < sx; x++ {
			chr := a.GetCharOfElementAtCoords(x, y)
			setcolorForRune(chr)
			cw.PutChar(chr, x+sx*(g.roomW+1)+2, y)
		}
	}
	cw.SetColor(cw.BEIGE, cw.BLACK)
	cw.PutString(fmt.Sprintf("PATTERN SELECTED: #%d  ", desiredPNum), sx*(g.roomW+1)+2, sy+2)
	cw.PutString(fmt.Sprintf("PATTERN USED: #%d  ", pattNum), sx*(g.roomW+1)+2, sy+3)
	cw.PutString(fmt.Sprintf("FILE: %s  ", fName), sx*(g.roomW+1)+2, sy+4)
	// cw.PutString(fmt.Sprintf("NAME: %s  ", pName), sx*(g.roomW+1)+2, sy+5)
	cw.PutString(fmt.Sprintf("%dx%d nodes", W, H), sx*(g.roomW+1)+2, sy+6)
	if restarts > maxDesiredRestarts {
		cw.SetColor(cw.BLACK, cw.RED)
	}
	cw.PutString(fmt.Sprintf("Gen restarts: %d", restarts), sx*(g.roomW+1)+2, sy+7)
	cw.SetColor(cw.BEIGE, cw.BLACK)
	cw.PutString(fmt.Sprintf("Room size: %dx%d", g.roomW, g.roomH), sx*(g.roomW+1)+2, sy+8)
	if rand {
		cw.PutString("Random paths", sx*(g.roomW+1)+2, sy+9)
	} else {
		cw.PutString("Shortest paths", sx*(g.roomW+1)+2, sy+9)
	}
}

func (g *tiledMapVisualiser) drawLevel(level *[][]layout_tiler.Tile, sx, sy int) {
	for x := 0; x < len(*level); x++ {
		for y := 0; y < len((*level)[x]); y++ {
			chr := (*level)[x][y].GetChar()
			setcolorForRune(chr)

			code := (*level)[x][y].Code
			lockId := (*level)[x][y].LockId
			if code == layout_tiler.TILE_DOOR {
				if lockId != 0 {
					chr = rune(strconv.Itoa(lockId)[0])
					cw.SetColor(cw.BLACK, cw.DARK_MAGENTA)
				}
			}
			if code == layout_tiler.TILE_KEY_PLACE {
				chr = rune(strconv.Itoa(lockId)[0])
				cw.SetColor(cw.DARK_MAGENTA, cw.BLACK)
			}
			cw.PutChar(chr, sx+x, sy+y)
			cw.SetColor(cw.WHITE, cw.BLACK)
		}
	}
}
