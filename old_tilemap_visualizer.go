package main

import (
	"CyclicDungeonGenerator/layout_generation"
	"CyclicDungeonGenerator/deprecated/layout_to_tilemap"
	"fmt"
	cw "CyclicDungeonGenerator/console_wrapper"
	"CyclicDungeonGenerator/random"
)

type tmv struct {}

func (g *tmv) doTilemapVisualization() {
	key := "none"
	desiredPatternNum := -1
	rnd := random.FibRandom{}
	rnd.InitDefault()
	layout_to_tilemap.Random = &rnd
	parser := layout_generation.PatternParser{}
	filenames := parser.ListPatternFilenamesInPath("patterns/")

	for key != "ESCAPE" {
		cw.Clear_console()
		pattNum := rnd.Rand(len(filenames))
		if desiredPatternNum != -1 {
			pattNum = desiredPatternNum
		}
		gen := layout_generation.InitCyclicGenerator(true, W, H, -1)
		generatedMap, genRestarts := gen.GenerateLayout(parser.ParsePatternFile(filenames[pattNum]))

		if generatedMap == nil {
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
			g.putTileMap(generatedMap)
			// putMiniMapAndPatternNumberAndNumberOfTries(generatedMap, pattNum, desiredPatternNum, genRestarts)
		}
		cw.Flush_console()
	keyread:
		for {
			key = cw.ReadKey()
			switch key {
			case "=":
				if desiredPatternNum < len(filenames)-1 {
					desiredPatternNum++
				}
				break keyread
			case "-":
				if desiredPatternNum > -1 {
					desiredPatternNum--
				}
				break keyread
			case " ", "ESCAPE":
				break keyread
			}
		}
	}
}

func (g *tmv) putTileMap(a *layout_generation.LayoutMap) {
	g.putTileArray(layout_to_tilemap.TransformLayoutToTileMap(a), 0, 0)
	rw, rh := a.GetSize()
	for rx := 0; rx < rw; rx++ {
		for ry := 0; ry < rh; ry++ {
			node := a.GetElement(rx, ry)
			conns := node.GetAllConnectionsCoords()
			if len(conns) > 0 {
				roomSize := 11 // temp
				cw.SetFgColor(cw.GREEN)
				if node.IsNode() {
					name := node.GetName()
					namelen := len(name)
					offset := roomSize / 2 - namelen / 2
					cw.PutString(name, rx*roomSize + offset, ry*roomSize+roomSize/2)
				}
			}
		}
	}
}

func (g *tmv) putTileArray(arr *[][]layout_to_tilemap.Tile, sx, sy int) {
	for x :=0; x <len(*arr); x++{
		for y :=0; y <len((*arr)[x]); y++ {
			chr := (*arr)[x][y].Char
			setcolorForRune(chr)
			cw.PutChar(chr, sx+x, sy+y)
		}
	}
}
