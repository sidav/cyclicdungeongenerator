package layout_generation

type nodeCell struct {
	nodeName string
	nodeTag  string // Locked? Key-containing? Anything.
}

func (n *nodeCell) AddTag(tag string) { // TODO: make the tags an array
	n.nodeTag = tag
}
