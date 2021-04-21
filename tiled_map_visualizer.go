package main

import (
	cw "CyclicDungeonGenerator/console_wrapper"
	"CyclicDungeonGenerator/layout_generation"
	"CyclicDungeonGenerator/layout_to_tiled_map"
	"CyclicDungeonGenerator/random"
	"fmt"
	"strconv"
)

type vis struct {
	roomW, roomH                int
	drawRoomNames, drawRoomTags bool
}

func (g *vis) doTilemapVisualization() {
	g.roomW, g.roomH = 5, 3
	key := "none"
	desiredPatternNum := -1
	rnd := random.FibRandom{}
	rnd.InitDefault()
	parser := layout_generation.PatternParser{}
	filenames := parser.ListPatternFilenamesInPath("patterns/")

	for key != "ESCAPE" {
		cw.Clear_console()
		pattNum := rnd.Rand(len(filenames))
		if desiredPatternNum != -1 {
			pattNum = desiredPatternNum
		}
		gen := layout_generation.InitCyclicGenerator(true, W, H, -1)
		gen.TriesForPattern = 100
		patt := parser.ParsePatternFile(filenames[pattNum])
		generatedMap, genRestarts := gen.GenerateLayout(patt)

		if generatedMap == nil {
			key = ""
			cw.PutString(":(", 0, 0)
			cw.PutString(fmt.Sprintf("Generation failed even after %d restarts, pattern #%d", genRestarts, pattNum), 0, 1)
			cw.PutString("Press ENTER to generate again or ESCAPE to exit.", 0, 2)
			cw.Flush_console()
			for key != "ESCAPE" && key != "ENTER" {
				key = cw.ReadKey()
			}
			if key == "ENTER" {
				continue
			} else {
				break
			}
		} else {
			g.convertLayoutToLevel(&rnd, generatedMap)
			g.putInfo(generatedMap, pattNum, desiredPatternNum, patt.Filename, patt.Name, genRestarts, gen.RandomizePath)
			// putMiniMapAndPatternNumberAndNumberOfTries(generatedMap, pattNum, desiredPatternNum, genRestarts)
		}
		cw.Flush_console()
	keyread:
		for {
			key = cw.ReadKey()
			switch key {
			case "b":
				g.roomW++
				g.roomH++
				break keyread
			case "s":
				g.roomW--
				g.roomH--
				break keyread
			case "=", "+", "RIGHT":
				if desiredPatternNum < len(filenames)-1 {
					desiredPatternNum++
				}
				break keyread
			case "-", "LEFT":
				if desiredPatternNum > -1 {
					desiredPatternNum--
				}
				break keyread
			case " ", "TAB", "ESCAPE":
				break keyread
			}
		}
	}
}

func (g *vis) putInfo(a *layout_generation.LayoutMap, pattNum, desiredPNum int, fName, pName string, restarts int, rand bool) {
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
	cw.PutString(fmt.Sprintf("NAME: %s  ", pName), sx*(g.roomW+1)+2, sy+5)
	cw.PutString(fmt.Sprintf("Gen restarts: %d", restarts), sx*(g.roomW+1)+2, sy+6)
	cw.PutString(fmt.Sprintf("Room size: %dx%d", g.roomW, g.roomH), sx*(g.roomW+1)+2, sy+7)
	if rand {
		cw.PutString("Random paths", sx*(g.roomW+1)+2, sy+8)
	} else {
		cw.PutString("Shortest paths", sx*(g.roomW+1)+2, sy+8)
	}
}

func (g *vis) convertLayoutToLevel(rnd *random.FibRandom, layout *layout_generation.LayoutMap) {
	cw.Clear_console()
	ltl := layout_to_tiled_map.LayoutToLevel{}
	ltl.Init(rnd, g.roomW, g.roomH)
	ltl.CAConnectionChance = 100
	ltl.CARoomChance = 15
	ltl.ProcessLayout(layout, "layout_to_tiled_map/submaps/")
	g.drawLevel(&ltl.TileMap, 0, 0)
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

func (g *vis) drawLevel(level *[][]layout_to_tiled_map.Tile, sx, sy int) {
	for x := 0; x < len(*level); x++ {
		for y := 0; y < len((*level)[x]); y++ {
			chr := (*level)[x][y].GetChar()
			setcolorForRune(chr)

			code := (*level)[x][y].Code
			lockId := (*level)[x][y].LockId
			if code == layout_to_tiled_map.TILE_DOOR {
				if lockId != 0 {
					chr = rune(strconv.Itoa(lockId)[0])
					cw.SetColor(cw.BLACK, cw.DARK_MAGENTA)
				}
			}
			if code == layout_to_tiled_map.TILE_KEY_PLACE {
				chr = rune(strconv.Itoa(lockId)[0])
				cw.SetColor(cw.DARK_MAGENTA, cw.BLACK)
			}
			cw.PutChar(chr, sx+x, sy+y)
			cw.SetColor(cw.WHITE, cw.BLACK)
		}
	}
}
