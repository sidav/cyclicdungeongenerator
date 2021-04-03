package layout_generation

type nodeCell struct {
	nodeName string
	nodeStatus string // Locked? Key-containing? Anything.
}

func (n *nodeCell) AddStatus(status string) { // TODO: make the status an array
	n.nodeStatus = status
}
