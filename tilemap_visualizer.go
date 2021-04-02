package main

import (
	"CyclicDungeonGenerator/layout_generation"
	"CyclicDungeonGenerator/layout_to_tiled"
	"fmt"
	cw "CyclicDungeonGenerator/console_wrapper"
	rnd "CyclicDungeonGenerator/random"
)

type tmv struct {}

func (g *tmv) doTilemapVisualization() {
	key := "none"
	desiredPatternNum := -1

	for key != "ESCAPE" {
		cw.Clear_console()
		pattNum := rnd.Random(layout_generation.GetTotalPatternsNumber())
		if desiredPatternNum != -1 {
			pattNum = desiredPatternNum
		}
		generatedMap, genRestarts := layout_generation.Generate(pattNum, W, H)

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
			g.putTileMap(generatedMap)
			// putMiniMapAndPatternNumberAndNumberOfTries(generatedMap, pattNum, desiredPatternNum, genRestarts)
		}
		cw.Flush_console()
	keyread:
		for {
			key = cw.ReadKey()
			switch key {
			case "=":
				if desiredPatternNum < layout_generation.GetTotalPatternsNumber()-1 {
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
	g.putTileArray(layout_to_tiled.TransformLayoutToTileMap(a), 0, 0)
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

func (g *tmv) putTileArray(arr *[][]layout_to_tiled.Tile, sx, sy int) {
	for x :=0; x <len(*arr); x++{
		for y :=0; y <len((*arr)[x]); y++ {
			chr := (*arr)[x][y].Char
			setcolorForRune(chr)
			cw.PutChar(chr, sx+x, sy+y)
		}
	}
}
