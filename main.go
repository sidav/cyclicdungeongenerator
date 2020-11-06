package main

import (
	rnd "github.com/sidav/golibrl/random"
	cw "github.com/sidav/golibrl/console/tcell_console"
)

const (
	W = 5
	H = 5
)

func main() {
	rnd.Randomize()

	// layout_generation.Benchmark(-1, false, true)

	// cw.Init_console("CDG", cw.TCellRenderer)
	cw.Init_console()
	defer cw.Close_console()

	// doLayoutVisualization()
	//tmv := tmv{}
	//tmv.doTilemapVisualization()
	gen := generatedVisualizer{}
	gen.doGeneratedVisualization()
}
