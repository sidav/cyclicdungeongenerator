package main

import (
	cw "CyclicDungeonGenerator/console_wrapper"
	"CyclicDungeonGenerator/layout_generation"
)

var (
	W = 5
	H = 5
)

func main() {

	const bench = false
	if bench {
		bnch := layout_generation.Benchmark{
			LayoutWidth: W,
			LayoutHeight: H,

			BenchLoopsForPattern:            10000,
			TriesForPattern:                 25,
			CheckRandomPaths:                true,
			CheckShortestPaths:              true,
			TestUniquity:                    true,
			GenerateAndConsiderGarbageNodes: false,
		}
		bnch.Benchmark("patterns/")
	}

	// cw.Init_console("CDG", cw.TCellRenderer)
	cw.Init_console()
	defer cw.Close_console()

	// doLayoutVisualization()
	//tmv := tmv{}
	//tmv.doTilemapVisualization()
	vis := vis{}
	vis.doTilemapVisualization()
}
