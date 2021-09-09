package layout_generation

const (
	ACTION_PLACE_NODE_AT_EMPTY = iota
	ACTION_PLACE_NODE_NEAR_PATH
	ACTION_PLACE_RANDOM_CONNECTED_NODES     // For additional "garbage" dead end nodes. Can be placed connected to anything: a node, a path...
	ACTION_FILL_WITH_RANDOM_CONNECTED_NODES // For additional "garbage" dead end nodes. Can be placed connected to anything: a node, a path...
	ACTION_PLACE_PATH_FROM_TO
	ACTION_PLACE_OBSTACLE_IN_CENTER
	ACTION_PLACE_OBSTACLE_AT_COORDS
	ACTION_CLEAR_OBSTACLES
	ACTION_PLACE_NODE_AT_PATH
	ACTION_SET_NODE_TAGS
	ACTION_LOCK_PATH
	ACTION_SET_NODE_CONNECTIONS_LOCKED
	ACTION_GROW_ALL_NODES
)

type patternStep struct {
	pattern                      *pattern
	actionType                   int
	fromX, fromY, toX, toY       int    // for coordinate ranges
	minEmptyCellsNear            int    // for nodes placement.
	maxNodeSize                  int    // for random "node grow"
	countFrom, countTo           int    // for random ranges
	nameOfNode, nameFrom, nameTo string // for node names
	obstacleRadius               int    // for centered non-random obstacle
	pathNumber                   int    // for... path numbering, I guess o_O
	allowPlaceNearNode           bool   // for ACTION_PLACE_NODE_NEAR_PATH
	tags                         []string // for ACTION_SET_NODE_TAGS
	lockNumber                   int    // for various locks
	allowCrossPaths              bool   // for allowing the new path to cross already existing one
	instructionText              string // is written only if the parser is set so
}
