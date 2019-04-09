package main

import (
	"CyclicDungeonGenerator/layout_generation"
	"CyclicDungeonGenerator/layout_to_tiled"
	"fmt"
	cw "github.com/sidav/golibrl/console/tcell_console"
	rnd "github.com/sidav/golibrl/random"
)

func doTilemapVisualization() {
	key := "none"
	desiredPatternNum := 0

	for key != "ESCAPE" {
		cw.Clear_console()
		pattNum := rnd.Random(layout_generation.GetTotalPatternsNumber())
		if desiredPatternNum != -1 {
			pattNum = desiredPatternNum
		}
		generatedMap, genRestarts := layout_generation.Generate(pattNum)

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
			putTileMap(generatedMap)
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

func putTileMap(a *layout_generation.LayoutMap) {
	rw, rh := a.GetSize()
	for rx := 0; rx < rw; rx++ {
		for ry := 0; ry < rh; ry++ {
			node := a.GetElement(rx, ry)
			conns := node.GetAllConnectionsCoords()
			roomStrs := layout_to_tiled.GetRoomByNodeConnections(&conns)
			roomSize := len(*roomStrs) - 1
			putStringArray(roomStrs, ry*roomSize, rx*roomSize)
		}
	}
}

func putStringArray(arr *[]string, sx, sy int) {
	for x :=0; x <len(*arr); x++{
		for y :=0; y <len((*arr)[x]); y++ {
			chr := rune((*arr)[x][y])
			setcolorForRune(chr)
			cw.PutChar(chr, sy+y, sx+x)
		}
	}
}
