package main

import (
	cw "CyclicDungeonGenerator/console_wrapper"
	"CyclicDungeonGenerator/layout_generation"
	"CyclicDungeonGenerator/layout_to_tiled_map"
	"CyclicDungeonGenerator/random"
	"fmt"
)

type vis struct {
	roomW, roomH int
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
			cw.PutString(":(", 0, 0)
			cw.PutString(fmt.Sprintf("Generation failed even after %d restarts, pattern #%d", genRestarts, pattNum), 0, 1)
			cw.PutString("Press ENTER to generate again or ESCAPE to exit.", 0, 2)
			cw.Flush_console()
			for key != "ESCAPE" && key != "ENTER" {
				key = cw.ReadKey()
			}
			continue
		} else {
			g.putTileMap(&rnd, generatedMap)
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

func (g *vis) putTileMap(rnd *random.FibRandom, layout *layout_generation.LayoutMap) {
	cw.Clear_console()
	ltl := layout_to_tiled_map.LayoutToLevel{}
	ltl.Init(rnd, g.roomW, g.roomH)
	ltl.CAConnectionChance = 100
	ltl.CARoomChance = 15
	g.putTileArray(ltl.MakeCharmap(layout, "layout_to_tiled_map/submaps/"), 0, 0)
	rw, rh := layout.GetSize()
	for rx := 0; rx < rw; rx++ {
		for ry := 0; ry < rh; ry++ {
			node := layout.GetElement(rx, ry)
			conns := node.GetAllConnectionsCoords()
			if len(conns) > 0 {
				cw.SetFgColor(cw.GREEN)
				if node.IsNode() {
					name := node.GetName()
					strlen := len(name)
					offset := (g.roomW+1) / 2 - strlen/ 2
					cw.PutString(name, rx*(g.roomW+1)  + offset, ry*(g.roomH+1)+(g.roomH+1)/2)
					tags := node.GetTags()
					strlen = len(tags)
					offset = (g.roomW+1) / 2 - strlen/ 2
					cw.PutString(tags, rx*(g.roomW+1)  + offset, ry*(g.roomH+1)+(g.roomH+1)/2+1)
				}
			}
		}
	}
}

func (g *vis) putTileArray(arr *[][]rune, sx, sy int) {
	for x :=0; x <len(*arr); x++{
		for y :=0; y <len((*arr)[x]); y++ {
			chr := (*arr)[x][y]
			setcolorForRune(chr)
			cw.PutChar(chr, sx+x, sy+y)
			cw.SetColor(cw.WHITE, cw.BLACK)
		}
	}
}
