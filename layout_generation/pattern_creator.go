package layout_generation

import (
	rnd "CyclicDungeonGenerator/random"
)

type pattern struct {
	name string
	instructions []patternStep
}

var patternsArray = []pattern{
	{
		name: "Single key",
		instructions: []patternStep{
			// randomly place beginning and end
			{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "STA", minEmptyCellsNear: 4},
			{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "FIN", minEmptyCellsNear: 4},
			// place big obstacle in center and some random obstacles for path to be less straight
			// {actionType: ACTION_PLACE_OBSTACLE_IN_CENTER, obstacleRadius: 1},
			// {actionType: ACTION_PLACE_RANDOM_OBSTACLES, countFrom: 2, countTo: 3},
			// draw two paths
			{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "STA", nameTo: "FIN", pathNumber: 1},
			{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "STA", nameTo: "FIN", pathNumber: 2},
			// clear temp obstacles
			// {actionType: ACTION_CLEAR_OBSTACLES},
			//place new node and a path to the FIN
			{actionType: ACTION_PLACE_NODE_NEAR_PATH, pathNumber: 1, nameOfNode: "NDE"},
			{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "NDE", nameTo: "FIN", pathNumber: 3},
			// place garbage nodes
			{actionType: ACTION_PLACE_RANDOM_CONNECTED_NODES, countFrom: 3, countTo: 6},
			{actionType: ACTION_SET_NODE_STATUS, nameOfNode: "NDE", status: "KEY"},
			{actionType: ACTION_SET_NODE_CONNECTION_LOCKED_FROM_PATH, nameOfNode: "FIN", pathNumber: 1, lockNumber: 0},
		},
	},
	// pattern 1, more complicated pattern.
	{
		name: "No keys, one loop",
		instructions: []patternStep{
			// randomly place beginning and end
			{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "STA", fx: 1, fy: 1, tx: -2, ty: -2, minEmptyCellsNear: 4},
			{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "FIN", fx: 1, fy: 1, tx: -2, ty: -2, minEmptyCellsNear: 4},
			// place big obstacle in center and some random obstacles for path to be less straight
			// {actionType: ACTION_PLACE_OBSTACLE_IN_CENTER, obstacleRadius: 1},
			// {actionType: ACTION_PLACE_RANDOM_OBSTACLES, countFrom: 2, countTo: 3},
			// draw two paths
			{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "STA", nameTo: "FIN", pathNumber: 1},
			// clear temp obstacles
			// {actionType: ACTION_CLEAR_OBSTACLES},
			//place new node and a path to the FIN
			{actionType: ACTION_PLACE_NODE_NEAR_PATH, pathNumber: 1, nameOfNode: "ND1"},
			{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "ND2", minEmptyCellsNear: 2},
			{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "ND1", nameTo: "FIN", pathNumber: 2},
			{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "ND1", nameTo: "ND2", pathNumber: 3},
			{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "ND2", nameTo: "FIN", pathNumber: 4},
			// place garbage nodes
			{actionType: ACTION_PLACE_RANDOM_CONNECTED_NODES, countFrom: 2, countTo: 5},
		},
	},
	// pattern 2, by Quosteeque.
	{
		name: "No keys, two loops",
		instructions: []patternStep{
			{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "STA", fx: 1, fy: 1, tx: -2, ty: -2, minEmptyCellsNear: 4},
			{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "FIN", fx: 1, fy: 1, tx: -2, ty: -2, minEmptyCellsNear: 4},
			{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "ND1", fx: 1, fy: 1, tx: -2, ty: -2, minEmptyCellsNear: 0},
			{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "ND2", fx: 1, fy: 1, tx: -2, ty: -2, minEmptyCellsNear: 0},
			{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "STA", nameTo: "FIN", pathNumber: 1},
			{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "STA", nameTo: "ND1", pathNumber: 2},
			{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "STA", nameTo: "ND2", pathNumber: 3},
			{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "ND1", nameTo: "FIN", pathNumber: 4},
			{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "ND2", nameTo: "FIN", pathNumber: 5},
			{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "ND1", nameTo: "ND2", pathNumber: 6},
			// place garbage nodes
			{actionType: ACTION_PLACE_RANDOM_CONNECTED_NODES, countFrom: 2, countTo: 7},
		},
	},
	// pattern 3, by Quosteeque too.
	{
		name: "No keys, four loops",
		instructions: []patternStep{
			{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "STA", fx: 1, fy: 1, tx: -2, ty: -2, minEmptyCellsNear: 4},
			{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "FIN", fx: 1, fy: 1, tx: -2, ty: -2, minEmptyCellsNear: 4},
			{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "ND1", fx: 1, fy: 1, tx: -2, ty: -2, minEmptyCellsNear: 0},
			{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "ND2", fx: 1, fy: 1, tx: -2, ty: -2, minEmptyCellsNear: 0},
			{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "STA", nameTo: "FIN", pathNumber: 1},
			{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "STA", nameTo: "ND1", pathNumber: 2},
			{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "STA", nameTo: "ND2", pathNumber: 3},
			{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "ND1", nameTo: "ND2", pathNumber: 4},
			{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "ND1", nameTo: "FIN", pathNumber: 5},
			{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "ND2", nameTo: "FIN", pathNumber: 6},
			// place garbage nodes
			{actionType: ACTION_FILL_WITH_RANDOM_CONNECTED_NODES},
		},
	},
	// pattern 4, mine.
	{
		name: "Red key required",
		instructions: []patternStep{
			{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "STA", fx: 1, fy: 1, tx: -2, ty: -2, minEmptyCellsNear: 4},
			{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "FIN", fx: 1, fy: 1, tx: -2, ty: -2, minEmptyCellsNear: 4},
			{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "ND1", fx: 1, fy: 1, tx: -2, ty: -2, minEmptyCellsNear: 1},
			{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "ND2", fx: 1, fy: 1, tx: -2, ty: -2, minEmptyCellsNear: 1},
			{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "STA", nameTo: "FIN", pathNumber: 1},
			{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "STA", nameTo: "ND1", pathNumber: 2},
			{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "STA", nameTo: "ND2", pathNumber: 3},
			{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "ND1", nameTo: "ND2", pathNumber: 4},
			{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "ND1", nameTo: "FIN", pathNumber: 5},

			{actionType: ACTION_SET_NODE_STATUS, nameOfNode: "ND2", status: "KEY"},
			{actionType: ACTION_SET_NODE_CONNECTION_LOCKED_FROM_PATH, nameOfNode: "FIN", pathNumber: 1, lockNumber: 0},
			{actionType: ACTION_SET_NODE_CONNECTION_LOCKED_FROM_PATH, nameOfNode: "FIN", pathNumber: 5, lockNumber: 0},

			// place garbage nodes
			{actionType: ACTION_FILL_WITH_RANDOM_CONNECTED_NODES},
		},
	},
	// pattern 5.
	{
		name: "Red key required 2",
		instructions: []patternStep{
			{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "STA", fx: 1, fy: 1, tx: -2, ty: -2, minEmptyCellsNear: 4},
			{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "FIN", fx: 1, fy: 1, tx: -2, ty: -2, minEmptyCellsNear: 4},
			{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "ND1", fx: 1, fy: 1, tx: -2, ty: -2, minEmptyCellsNear: 1},
			{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "ND2", fx: 1, fy: 1, tx: -2, ty: -2, minEmptyCellsNear: 1},
			{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "ND3", fx: 1, fy: 1, tx: -2, ty: -2, minEmptyCellsNear: 1},
			{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "STA", nameTo: "ND1", pathNumber: 1},
			{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "STA", nameTo: "ND2", pathNumber: 2},
			{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "STA", nameTo: "ND3", pathNumber: 3},
			{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "ND1", nameTo: "ND2", pathNumber: 4},
			{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "ND1", nameTo: "ND3", pathNumber: 5},
			{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "ND3", nameTo: "FIN", pathNumber: 6},

			{actionType: ACTION_SET_NODE_STATUS, nameOfNode: "ND2", status: "KEY"},
			{actionType: ACTION_SET_NODE_CONNECTION_LOCKED_FROM_PATH, nameOfNode: "FIN", pathNumber: 1, lockNumber: 0},
			{actionType: ACTION_SET_NODE_CONNECTION_LOCKED_FROM_PATH, nameOfNode: "ND3", pathNumber: 3, lockNumber: 0},
			{actionType: ACTION_SET_NODE_CONNECTION_LOCKED_FROM_PATH, nameOfNode: "ND3", pathNumber: 5, lockNumber: 0},

			// place garbage nodes
			{actionType: ACTION_FILL_WITH_RANDOM_CONNECTED_NODES},
		},
	},
	// pattern 6, two-keyed
	{
		name: "Two keys",
		instructions: []patternStep{
			{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "STA", minEmptyCellsNear: 4},
			{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "ND1", fx: 1, fy: 1, tx: -2, ty: -2, minEmptyCellsNear: 1},
			{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "ND2", fx: 1, fy: 1, tx: -2, ty: -2, minEmptyCellsNear: 1},
			{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "ND3", fx: 1, fy: 1, tx: -2, ty: -2, minEmptyCellsNear: 1},
			{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "ND4", minEmptyCellsNear: 1},
			{actionType: ACTION_PLACE_NODE_AT_EMPTY, nameOfNode: "FIN", minEmptyCellsNear: 0},
			{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "STA", nameTo: "ND1", pathNumber: 1},
			{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "STA", nameTo: "ND2", pathNumber: 2},
			{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "ND1", nameTo: "ND3", pathNumber: 3},
			{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "ND2", nameTo: "ND3", pathNumber: 4},
			{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "ND3", nameTo: "ND4", pathNumber: 5},
			{actionType: ACTION_PLACE_PATH_FROM_TO, nameFrom: "ND4", nameTo: "FIN", pathNumber: 6},

			{actionType: ACTION_SET_NODE_STATUS, nameOfNode: "ND1", status: "KY1"},
			{actionType: ACTION_SET_NODE_STATUS, nameOfNode: "ND2", status: "KY2"},
			{actionType: ACTION_SET_NODE_CONNECTION_LOCKED_FROM_PATH, nameOfNode: "ND4", pathNumber: 5, lockNumber: 0},
			{actionType: ACTION_SET_NODE_CONNECTION_LOCKED_FROM_PATH, nameOfNode: "FIN", pathNumber: 6, lockNumber: 1},

			// place garbage nodes
			{actionType: ACTION_FILL_WITH_RANDOM_CONNECTED_NODES},
		},
	},
}

func getPattern(patternNumber int) *pattern {
	return &patternsArray[patternNumber%len(patternsArray)]
}

func getRandomPatternNumber() int {
	patternNumber := rnd.Random(len(patternsArray))
	return patternNumber
}

func GetTotalPatternsNumber() int {
	return len(patternsArray)
}
