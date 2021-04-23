package layout_generation

import "fmt"

type element struct {
	// it's a room or a tile occupied with interconnection.
	pathInfo    *pathCell
	nodeInfo    *nodeCell
	isObstacle  bool // for temp obstacles
	connections map[string]*connection
}

func (e *element) isPartOfAPath() bool {
	return e.pathInfo != nil
}

func (e *element) isEmpty() bool {
	return e.pathInfo == nil && e.nodeInfo == nil && !e.isObstacle
}

func (e *element) isPathOrEmpty() bool {
	return e.nodeInfo == nil && !e.isObstacle
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
		panic(fmt.Sprintf("ERROR PLACING CONNECTION: (%d,%d)\n", x, y))
		return
	}
	e.connections[direction] = c
}

func (e *element) GetConnectionByCoords(x, y int) *connection {
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

// exported

func (e *element) IsNode() bool {
	return e.nodeInfo != nil
}

func (e *element) GetName() string {
	return e.nodeInfo.nodeName
}

func (e *element) GetTags() string {
	return e.nodeInfo.nodeTag
}

func (e *element) GetAllConnectionsCoords() [][]int {
	arr := make([][]int, 0)
	if e.connections["north"] != nil {
		arr = append(arr, []int{0, -1})
	}
	if e.connections["south"] != nil {
		arr = append(arr, []int{0, 1})
	}
	if e.connections["west"] != nil {
		arr = append(arr, []int{-1, 0})
	}
	if e.connections["east"] != nil {
		arr = append(arr, []int{1, 0})
	}
	return arr
}
