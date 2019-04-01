package layout_generation

var patterns_array = [][]*patternStep{
	// first pattern.
	{
		// randomly place beginning and end
		&patternStep{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "Start", fx: 1, fy: 1, tx: -2, ty: -2},
		&patternStep{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "Finish", fx: 1, fy: 1, tx: -2, ty: -2},
		// place big obstacle in center and some random obstacles for path to be less straight
		&patternStep{actionType: ACTION_PLACE_OBSTACLE_IN_CENTER, obstacleRadius: 1},
		&patternStep{actionType: ACTION_PLACE_RANDOM_OBSTACLES, countFrom: 2, countTo: 3},
		// draw two paths
		&patternStep{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "Start", nameTo: "Finish", pathNumber: 1},
		&patternStep{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "Start", nameTo: "Finish", pathNumber: 2},
		// clear temp obstacles
		&patternStep{actionType: ACTION_CLEAR_OBSTACLES},
		//place new node and a path to the finish
		&patternStep{actionType: ACTION_PLACE_NODE_NEAR_PATH, pathNumber: 1, nameOfNode: "Node1"},
		&patternStep{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "Node1", nameTo: "Finish", pathNumber: 3},
		// place garbage node
		&patternStep{actionType: ACTION_PLACE_NODE_NEAR_PATH, pathNumber: -1, nameOfNode: "Mode2"},
	},
}

func getPattern(patternNumber int) []*patternStep {
	return patterns_array[patternNumber%len(patterns_array)]
}
