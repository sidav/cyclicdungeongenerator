package layout_generation

import (
	rnd "github.com/sidav/golibrl/random"
)

var patterns_array = [][]*patternStep{
	// first pattern.
	{
		// randomly place beginning and end
		&patternStep{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "STA", fx: 1, fy: 1, tx: -2, ty: -2},
		&patternStep{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "FIN", fx: 1, fy: 1, tx: -2, ty: -2},
		// place big obstacle in center and some random obstacles for path to be less straight
		&patternStep{actionType: ACTION_PLACE_OBSTACLE_IN_CENTER, obstacleRadius: 1},
		&patternStep{actionType: ACTION_PLACE_RANDOM_OBSTACLES, countFrom: 2, countTo: 3},
		// draw two paths
		&patternStep{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "STA", nameTo: "FIN", pathNumber: 1},
		&patternStep{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "STA", nameTo: "FIN", pathNumber: 2},
		// clear temp obstacles
		&patternStep{actionType: ACTION_CLEAR_OBSTACLES},
		//place new node and a path to the FIN
		&patternStep{actionType: ACTION_PLACE_NODE_NEAR_PATH, pathNumber: 1, nameOfNode: "NDE"},
		&patternStep{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "NDE", nameTo: "FIN", pathNumber: 3},
		// place garbage node
		&patternStep{actionType: ACTION_PLACE_NODE_NEAR_PATH, pathNumber: -1, nameOfNode: "   ", allowPlaceNearNode: true},
		&patternStep{actionType: ACTION_PLACE_NODE_NEAR_PATH, pathNumber: -1, nameOfNode: "   ", allowPlaceNearNode: true},
		&patternStep{actionType: ACTION_PLACE_NODE_NEAR_PATH, pathNumber: -1, nameOfNode: "   ", allowPlaceNearNode: true},
	},
	// Second, more complicated pattern.
	{
		// randomly place beginning and end
		&patternStep{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "STA", fx: 1, fy: 1, tx: -2, ty: -2},
		&patternStep{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "FIN", fx: 1, fy: 1, tx: -2, ty: -2},
		// place big obstacle in center and some random obstacles for path to be less straight
		&patternStep{actionType: ACTION_PLACE_OBSTACLE_IN_CENTER, obstacleRadius: 1},
		&patternStep{actionType: ACTION_PLACE_RANDOM_OBSTACLES, countFrom: 2, countTo: 3},
		// draw two paths
		&patternStep{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "STA", nameTo: "FIN", pathNumber: 1},
		// clear temp obstacles
		&patternStep{actionType: ACTION_CLEAR_OBSTACLES},
		//place new node and a path to the FIN
		&patternStep{actionType: ACTION_PLACE_NODE_NEAR_PATH, pathNumber: 1, nameOfNode: "ND1"},
		&patternStep{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "ND2", fx: 0, fy: 0, tx: -1, ty: -1},
		&patternStep{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "ND1", nameTo: "FIN", pathNumber: 2},
		&patternStep{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "ND1", nameTo: "ND2", pathNumber: 2},
		&patternStep{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "ND2", nameTo: "FIN", pathNumber: 3},
		// place garbage node
		&patternStep{actionType: ACTION_PLACE_NODE_NEAR_PATH, pathNumber: -1, nameOfNode: "   ", allowPlaceNearNode: true},
		&patternStep{actionType: ACTION_PLACE_NODE_NEAR_PATH, pathNumber: -1, nameOfNode: "   ", allowPlaceNearNode: true},
		&patternStep{actionType: ACTION_PLACE_NODE_NEAR_PATH, pathNumber: -1, nameOfNode: "   ", allowPlaceNearNode: true},
	},
}

func getPattern(patternNumber int) []*patternStep {
	return patterns_array[patternNumber%len(patterns_array)]
}

func getRandomPatternNumber() int {
	patternNumber := rnd.Random(len(patterns_array))
	return patternNumber
}

func getTotalPatternsNumber() int {
	return len(patterns_array)
}
