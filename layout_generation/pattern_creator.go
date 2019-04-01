package layout_generation

var patterns_array = [][]*patternStep{
	// first pattern.
	{
		// randomly place beginning and end
		&patternStep{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "Start", fx: 1, fy: 1, tx: 1, ty: 1},
		&patternStep{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "Finish", fx: -2, fy: -2, tx: -1, ty: -1},
		// place big obstacle in center and some random obstacles for path to be less straight
		&patternStep{actionType: ACTION_PLACE_OBSTACLE_IN_CENTER, obstacleRadius: 1},
		&patternStep{actionType: ACTION_PLACE_RANDOM_OBSTACLES, countFrom: 2, countTo: 3},
		// draw two paths
		&patternStep{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "Start", nameTo: "Finish"},
		&patternStep{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "Start", nameTo: "Finish"},
	},
}

func getPattern(patternNumber int) []*patternStep {
	return patterns_array[patternNumber%len(patterns_array)]
}
