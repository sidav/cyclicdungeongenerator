package main

import (
	"CyclicDungeonGenerator/layout_generation"
	rnd "CyclicDungeonGenerator/random"
	cw "CyclicDungeonGenerator/console_wrapper"
)

const (
	W = 5
	H = 5
)

func main() {
	const bench = false
	rnd.Randomize()

	if bench {
		bnch := layout_generation.Benchmark{
			BenchLoopsForPattern:            100000,
			TriesForPattern:                 25,
			CheckRandomPaths:                true,
			CheckShortestPaths:              true,
			TestUniquity:                    true,
			GenerateAndConsiderGarbageNodes: false,
		}
		bnch.Benchmark(-1)
	}

	// cw.Init_console("CDG", cw.TCellRenderer)
	cw.Init_console()
	defer cw.Close_console()

	doLayoutVisualization()
	// tmv := tmv{}
	// tmv.doTilemapVisualization()
	//gen := generatedVisualizer{}
	//gen.doGeneratedVisualization()
}
