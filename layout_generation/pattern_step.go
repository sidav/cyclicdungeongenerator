package layout_generation

const (
	ACTION_PLACE_NODE_AT_EMPTY              = iota
	ACTION_PLACE_NODE_NEAR_PATH             = iota
	ACTION_PLACE_RANDOM_CONNECTED_NODES     = iota // For additional "garbage" dead end nodes. Can be placed connected to anything: a node, a path...
	ACTION_FILL_WITH_RANDOM_CONNECTED_NODES = iota // For additional "garbage" dead end nodes. Can be placed connected to anything: a node, a path...
	ACTION_PLACE_PATH_FROM_TO               = iota
	ACTION_PLACE_OBSTACLE_IN_CENTER         = iota
	ACTION_PLACE_OBSTACLE_AT_COORDS         = iota
	// ACTION_PLACE_RANDOM_OBSTACLES               = iota // unneeded since there is random pathfinding
	ACTION_CLEAR_OBSTACLES                      = iota
	ACTION_PLACE_NODE_AT_PATH                   = iota
	ACTION_SET_NODE_TAGS                        = iota
	ACTION_SET_NODE_CONNECTION_LOCKED_FROM_PATH = iota
)

type patternStep struct {
	actionType                   int
	fromX, fromY, toX, toY       int    // for coordinate ranges
	minEmptyCellsNear            int    // for nodes placement.
	countFrom, countTo           int    // for random ranges
	nameOfNode, nameFrom, nameTo string // for node names
	obstacleRadius               int    // for centered non-random obstacle
	pathNumber                   int    // for... path numbering, I guess o_O
	allowPlaceNearNode           bool   // for ACTION_PLACE_NODE_NEAR_PATH
	tags                         string // for ACTION_SET_NODE_TAGS
	lockNumber                   int    // for various locks
	allowCrossPaths              bool   // for allowing the new path to cross already existing one
	instructionText              string // is written only if the parser is set so
}
