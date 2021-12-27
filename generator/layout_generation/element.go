package layout_generation

import "fmt"

type Element struct {
	// it's a room or a tile occupied with interconnection.
	pathInfo    *pathCell
	nodeInfo    *nodeCell
	isObstacle  bool // for temp obstacles
	connections map[string]*connection
}

func (e *Element) isPartOfAPath() bool {
	return e.pathInfo != nil
}

func (e *Element) IsEmpty() bool {
	return e.pathInfo == nil && e.nodeInfo == nil && !e.isObstacle
}

func (e *Element) IsPathOrEmpty() bool {
	return e.nodeInfo == nil && !e.isObstacle
}

func (e *Element) setConnectionByCoords(c *connection, x, y int) {
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

func (e *Element) GetConnectionByCoords(x, y int) *connection {
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

func (e *Element) IsNode() bool {
	return e.nodeInfo != nil
}

func (e *Element) GetName() string {
	return e.nodeInfo.nodeName
}

func (e *Element) GetTags() []string {
	return e.nodeInfo.nodeTags
}

func (e *Element) HasNoTags() bool {
	return len(e.nodeInfo.nodeTags) == 0
}

func (e *Element) HasTag(tag string) bool {
	for _, t := range e.nodeInfo.nodeTags {
		if t == tag {
			return true
		}
	}
	return false
}

func (e *Element) GetAllConnectionsCoords() [][]int {
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
