package layout_generation

type nodeCell struct {
	nodeName string
	nodeTag  string // Locked? Key-containing? Anything.
}

func (n *nodeCell) setTags(tags string) { // TODO: make the tags an array
	n.nodeTag = tags
}
