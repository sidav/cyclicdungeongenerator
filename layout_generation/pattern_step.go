package layout_generation

const (
	ACTION_NONE                     = iota
	ACTION_PLACE_NODE_AT_EMPTY      = iota
	ACTION_PLACE_NODE_AT_PATH       = iota
	ACTION_PLACE_NODE_NEAR_PATH     = iota
	ACTION_PLACE_PATH_FROM_TO       = iota
	ACTION_PLACE_OBSTACLE_IN_CENTER = iota
	ACTION_PLACE_RANDOM_OBSTACLES   = iota
	ACTION_CLEAR_OBSTACLES          = iota
)

type patternStep struct {
	actionType                   int
	fx, fy, tx, ty               int    // for coordinate ranges
	countFrom, countTo           int    // for random ranges
	nameOfNode, nameFrom, nameTo string // for node names
	obstacleRadius               int    // for centered non-random obstacle
	pathNumber                   int	// for... path numbering, I guess o_O
}
