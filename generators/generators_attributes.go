package generators

import "cyclicdungeongenerator/generators/layout_generation"

type layoutGenerationAttributes struct {
	LastGeneratedPatternName, LastGeneratedPatternFilename string
	RandomPaths bool
	MaxGenerationTries int
	patternParser layout_generation.PatternParser
}

type layoutTilingAttributes struct {
	ChanceToCaveAConnection, ChanceToCaveARoom int
}
