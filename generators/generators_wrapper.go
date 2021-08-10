package generators

import layout_generation "cyclicdungeongenerator/generators/layout_generation"

type GeneratorsWrapper struct {
	RandomPaths bool
	MaxGenerationTries int
	patternParser layout_generation.PatternParser
}

func InitGeneratorsWrapper() *GeneratorsWrapper {
	gw := &GeneratorsWrapper{}
	gw.patternParser = layout_generation.PatternParser{}
	return gw
}

func (gw *GeneratorsWrapper) ListPatternFilenamesInPath(path string) []string {
	return gw.patternParser.ListPatternFilenamesInPath(path)
}

func (gw *GeneratorsWrapper) GenerateLayout(W, H int, patternFilename string) (*layout_generation.LayoutMap, int){
	gen := layout_generation.InitCyclicGenerator(gw.RandomPaths, W, H, -1)
	gen.TriesForPattern = gw.MaxGenerationTries
	pattern := gw.patternParser.ParsePatternFile(patternFilename, true)
	return gen.GenerateLayout(pattern)
}

func (gw *GeneratorsWrapper) ConvertLayoutToTiledMap() {

}
