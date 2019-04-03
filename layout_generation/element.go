package layout_generation

import "fmt"

type element struct {
	// it's a room or a tile occupied with interconnection.
	pathInfo    *path_cell
	nodeInfo    *node_cell
	isObstacle  bool // for temp obstacles
	connections map[string]*connection
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

func (e *element) setConnectionByCoords(c *connection, x, y int) {
	direction := "wat?"
	if x == 0 && y == 1 {
		direction = "south"
	}
	if x == 0 && y == -1 {
		direction = "north"
	}
	if x == 1 && y == 0 {
		direction = "east"
	}
	if x == -1 && y == 0 {
		direction = "west"
	}
	if direction == "wat?" {
		fmt.Printf("ERROR PLACING CONNECTION: (%d,%d)\n", x, y)
		return
	}
	e.connections[direction] = c
}

func (e *element) getConnectionByCoords(x, y int) *connection {
	direction := "wat?"
	if x == 0 && y == 1 {
		direction = "south"
	}
	if x == 0 && y == -1 {
		direction = "north"
	}
	if x == 1 && y == 0 {
		direction = "east"
	}
	if x == -1 && y == 0 {
		direction = "west"
	}
	if direction == "wat?" {
		return nil
	}
	return e.connections[direction]
}

func (e *element) setAllConnectionsLockedForPath(pathNum int, lockNum int) {
	for _, v := range e.connections {
		if v != nil && v.pathNum == pathNum {
			v.isLocked = true
			v.lockNum = lockNum
		}
	}
}
