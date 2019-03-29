package layout_generation

type element struct {
	// it's a room or a tile occupied with interconnection.
	pathInfo    *path_cell
	nodeInfo    *node_cell
	isObstacle bool // for temp obstacles
}

func (e *element) isPartOfAPath() bool {
	return e.pathInfo != nil
}

func (e *element) isNode() bool {
	return e.nodeInfo != nil
}

func (e *element) isEmpty() bool {
	return e.pathInfo == nil && e.nodeInfo == nil && !e.isObstacle
}
