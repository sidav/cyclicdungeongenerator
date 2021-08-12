package main

import (
	cw "cyclicdungeongenerator/console_wrapper"
	"cyclicdungeongenerator/generator"
	"cyclicdungeongenerator/generator/layout_generation"
	"fmt"
	"os"
	"strconv"
)

var (
	W = 5
	H = 5
	genWrapper *generator.CyclicDungeonGenerator
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Printf(
			"Arguments: \n" +
				" -b num_loops tries_for_pattern map_w map_h: do benchmark \n" +
				" -l <layout_w> <layout_h>: generate and show layouts/tilemaps \n" +
				" -o <path_to_pattern>: show initial and optimized instructions order for pattern \n")
		return
	}

	if len(args) >= 3 {
		var err error
		W, err = strconv.Atoi(args[1])
		if err != nil {
			W = 5
		}
		H, err = strconv.Atoi(args[2])
		if err != nil {
			H = 5
		}
	}

	genWrapper = generator.InitGeneratorsWrapper()

	switch args[0] {
	case "-b", "b", "benchmark":
		bench(args)

	case "l", "-l", "layout", "t", "-t", "tilemap":
		cw.Init_console()
		defer cw.Close_console()
		visboth := visBoth{}
		visboth.do()

	case "o", "-o":
		if len(args) > 1 {
			pathToPattern := args[1]
			pp := layout_generation.PatternParser{WriteLinesInResult: true}
			pat := pp.ParsePatternFile(pathToPattern, false)
			po := layout_generation.PatternOptimizer{}
			po.ShowInitialAndOptimizedInstructionOrders(pat)
		} else {
			fmt.Println("Path to the pattern is missing.")
		}

	default:
		fmt.Printf(
			"Unknown argument \"%s\". Arguments: \n"+
				" -b num_loops tries_for_pattern map_w map_h: do benchmark \n"+
				" -l <layout_w> <layout_h>: generate and show layouts/tilemaps \n"+
				" -o <path_to_pattern>: show initial and optimized instructions order for pattern \n", args[0])
	}
}

func bench(args []string) {
	if len(args) < 5 {
		fmt.Printf("Too few arguments...\n")
		fmt.Printf("usage: -b num_loops tries_for_pattern map_w map_h <pattern file>\n")
		fmt.Printf("example: -b 10000 25 5 5 patterns/example_pattern.ptn \n")
		return
	}
	loops, err := strconv.Atoi(args[1])
	if err != nil {
		panic(err)
	}
	tries, err := strconv.Atoi(args[2])
	if err != nil {
		panic(err)
	}
	width, err := strconv.Atoi(args[3])
	if err != nil {
		panic(err)
	}
	height, err := strconv.Atoi(args[4])
	if err != nil {
		panic(err)
	}
	benchPatternsPath := "patterns/"
	if len(args) > 5 {
		benchPatternsPath = args[5]
	}
	bnch := layout_generation.Benchmark{
		LayoutWidth:  width,
		LayoutHeight: height,

		BenchLoopsForPattern:            loops,
		TriesForPattern:                 tries,
		CheckRandomPaths:                true,
		CheckShortestPaths:              false,
		TestUniquity:                    true,
		GenerateAndConsiderGarbageNodes: false,
	}
	bnch.Benchmark(benchPatternsPath)
}
