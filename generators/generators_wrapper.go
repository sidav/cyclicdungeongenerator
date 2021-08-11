package generators

import (
	"cyclicdungeongenerator/generators/layout_generation"
	"cyclicdungeongenerator/generators/layout_tiler"
	"cyclicdungeongenerator/random"
)

type GeneratorsWrapper struct {
	LayoutGenerationParams layoutGenerationAttributes
	TilingParams           layoutTilingAttributes
}

func InitGeneratorsWrapper() *GeneratorsWrapper {
	gw := &GeneratorsWrapper{}
	gw.LayoutGenerationParams.patternParser = layout_generation.PatternParser{}
	return gw
}

func (gw *GeneratorsWrapper) ListPatternFilenamesInPath(path string) []string {
	return gw.LayoutGenerationParams.patternParser.ListPatternFilenamesInPath(path)
}

func (gw *GeneratorsWrapper) GenerateLayout(W, H int, patternFilename string) (*layout_generation.LayoutMap, int) {
	gen := layout_generation.InitCyclicGenerator(gw.LayoutGenerationParams.RandomPaths, W, H, -1)
	gen.TriesForPattern = gw.LayoutGenerationParams.MaxGenerationTries
	pattern := gw.LayoutGenerationParams.patternParser.ParsePatternFile(patternFilename, true)
	return gen.GenerateLayout(pattern)
}

func (gw *GeneratorsWrapper) ConvertLayoutToTiledMap(
	rnd *random.FibRandom, layout *layout_generation.LayoutMap, roomW, roomH int, submapsDir string) [][]layout_tiler.Tile {
	ltl := layout_tiler.LayoutTiler{}
	ltl.Init(rnd, roomW, roomH)

	ltl.TagForEntry = "start"
	ltl.TagForExit = "finish"
	ltl.TagsForKeys = []string {"ky1", "ky2", "ky3"}

	ltl.CAConnectionChance = gw.TilingParams.ChanceToCaveAConnection
	ltl.CARoomChance = gw.TilingParams.ChanceToCaveARoom
	ltl.ProcessLayout(layout, submapsDir)
	return ltl.TileMap
}
