package layout_generation

const (
	ACTION_NONE                                 = iota
	ACTION_PLACE_NODE_AT_EMPTY                  = iota
	ACTION_PLACE_NODE_NEAR_PATH                 = iota
	ACTION_PLACE_RANDOM_CONNECTED_NODES         = iota // For additional "garbage" dead end nodes. Can be placed connected to anything: a node, a path...
	ACTION_PLACE_PATH_FROM_TO                   = iota
	ACTION_PLACE_OBSTACLE_IN_CENTER             = iota
	ACTION_PLACE_RANDOM_OBSTACLES               = iota
	ACTION_CLEAR_OBSTACLES                      = iota
	ACTION_PLACE_NODE_AT_PATH                   = iota
	ACTION_SET_NODE_STATUS                      = iota
	ACTION_SET_NODE_CONNECTION_LOCKED_FROM_PATH = iota
)

type patternStep struct {
	actionType                   int
	fx, fy, tx, ty               int    // for coordinate ranges
	countFrom, countTo           int    // for random ranges
	nameOfNode, nameFrom, nameTo string // for node names
	obstacleRadius               int    // for centered non-random obstacle
	pathNumber                   int    // for... path numbering, I guess o_O
	allowPlaceNearNode           bool   // for ACTION_PLACE_NODE_NEAR_PATH
	status                       string // for ACTION_SET_NODE_STATUS
	lockNumber                   int    // for vatious locks
}
