package layout_generation

import (
	rnd "github.com/sidav/golibrl/random"
)

var patterns_array = [][]*patternStep{
	// pattern 0
	{
		// randomly place beginning and end
		&patternStep{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "STA", fx: 1, fy: 1, tx: -2, ty: -2, minEmptyCellsNear: 4},
		&patternStep{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "FIN", fx: 1, fy: 1, tx: -2, ty: -2, minEmptyCellsNear: 4},
		// place big obstacle in center and some random obstacles for path to be less straight
		// &patternStep{actionType: ACTION_PLACE_OBSTACLE_IN_CENTER, obstacleRadius: 1},
		// &patternStep{actionType: ACTION_PLACE_RANDOM_OBSTACLES, countFrom: 2, countTo: 3},
		// draw two paths
		&patternStep{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "STA", nameTo: "FIN", pathNumber: 1},
		&patternStep{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "STA", nameTo: "FIN", pathNumber: 2},
		// clear temp obstacles
		// &patternStep{actionType: ACTION_CLEAR_OBSTACLES},
		//place new node and a path to the FIN
		&patternStep{actionType: ACTION_PLACE_NODE_NEAR_PATH, pathNumber: 1, nameOfNode: "NDE"},
		&patternStep{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "NDE", nameTo: "FIN", pathNumber: 3},
		// place garbage nodes
		&patternStep{actionType: ACTION_PLACE_RANDOM_CONNECTED_NODES, nameOfNode: "   ", countFrom: 3, countTo: 6},
		&patternStep{actionType: ACTION_SET_NODE_STATUS, nameOfNode: "NDE", status: "KEY"},
		&patternStep{actionType: ACTION_SET_NODE_CONNECTION_LOCKED_FROM_PATH, nameOfNode: "FIN", pathNumber: 1, lockNumber: 0},
	},
	// pattern 1, more complicated pattern.
	{
		// randomly place beginning and end
		&patternStep{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "STA", fx: 1, fy: 1, tx: -2, ty: -2, minEmptyCellsNear: 4},
		&patternStep{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "FIN", fx: 1, fy: 1, tx: -2, ty: -2, minEmptyCellsNear: 4},
		// place big obstacle in center and some random obstacles for path to be less straight
		&patternStep{actionType: ACTION_PLACE_OBSTACLE_IN_CENTER, obstacleRadius: 1},
		&patternStep{actionType: ACTION_PLACE_RANDOM_OBSTACLES, countFrom: 2, countTo: 3},
		// draw two paths
		&patternStep{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "STA", nameTo: "FIN", pathNumber: 1},
		// clear temp obstacles
		&patternStep{actionType: ACTION_CLEAR_OBSTACLES},
		//place new node and a path to the FIN
		&patternStep{actionType: ACTION_PLACE_NODE_NEAR_PATH, pathNumber: 1, nameOfNode: "ND1"},
		&patternStep{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "ND2", fx: 0, fy: 0, tx: -1, ty: -1, minEmptyCellsNear: 4},
		&patternStep{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "ND1", nameTo: "FIN", pathNumber: 2},
		&patternStep{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "ND1", nameTo: "ND2", pathNumber: 3},
		&patternStep{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "ND2", nameTo: "FIN", pathNumber: 4},
		// place garbage nodes
		&patternStep{actionType: ACTION_PLACE_RANDOM_CONNECTED_NODES, nameOfNode: "   ", countFrom: 2, countTo: 5},
	},
	// pattern 2, by Quosteeque.
	{

		&patternStep{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "STA", fx: 1, fy: 1, tx: -2, ty: -2, minEmptyCellsNear: 4},
		&patternStep{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "FIN", fx: 1, fy: 1, tx: -2, ty: -2, minEmptyCellsNear: 4},
		&patternStep{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "ND1", fx: 1, fy: 1, tx: -2, ty: -2, minEmptyCellsNear: 1},
		&patternStep{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "ND2", fx: 1, fy: 1, tx: -2, ty: -2, minEmptyCellsNear: 1},
		&patternStep{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "STA", nameTo: "FIN", pathNumber: 1},
		&patternStep{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "STA", nameTo: "ND1", pathNumber: 2},
		&patternStep{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "STA", nameTo: "ND2", pathNumber: 3},
		&patternStep{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "ND1", nameTo: "FIN", pathNumber: 4},
		&patternStep{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "ND2", nameTo: "FIN", pathNumber: 5},
		&patternStep{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "ND1", nameTo: "ND2", pathNumber: 6},
		// place garbage nodes
		&patternStep{actionType: ACTION_PLACE_RANDOM_CONNECTED_NODES, nameOfNode: "   ", countFrom: 2, countTo: 7},
	},
	// pattern 3, by Quosteeque too.
	{

		&patternStep{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "STA", fx: 1, fy: 1, tx: -2, ty: -2, minEmptyCellsNear: 4},
		&patternStep{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "FIN", fx: 1, fy: 1, tx: -2, ty: -2, minEmptyCellsNear: 4},
		&patternStep{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "ND1", fx: 1, fy: 1, tx: -2, ty: -2, minEmptyCellsNear: 1},
		&patternStep{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "ND2", fx: 1, fy: 1, tx: -2, ty: -2, minEmptyCellsNear: 1},
		&patternStep{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "STA", nameTo: "FIN", pathNumber: 1},
		&patternStep{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "STA", nameTo: "ND1", pathNumber: 2},
		&patternStep{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "STA", nameTo: "ND2", pathNumber: 3},
		&patternStep{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "ND1", nameTo: "ND2", pathNumber: 4},
		&patternStep{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "ND2", nameTo: "FIN", pathNumber: 5},
		&patternStep{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "ND2", nameTo: "FIN", pathNumber: 6},
		// place garbage nodes
		&patternStep{actionType: ACTION_PLACE_RANDOM_CONNECTED_NODES, nameOfNode: "   ", countFrom: 2, countTo: 7},
	},
}

func getPattern(patternNumber int) []*patternStep {
	return patterns_array[patternNumber%len(patterns_array)]
}

func getRandomPatternNumber() int {
	patternNumber := rnd.Random(len(patterns_array))
	return patternNumber
}

func GetTotalPatternsNumber() int {
	return len(patterns_array)
}
