package layout_generation

type node_cell struct {
	nodeName string
	nodeStatus string // Locked? Key-containing? Anything.
}

func (n *node_cell) AddStatus(status string) { // TODO: make the status an array
	n.nodeStatus = status
}
