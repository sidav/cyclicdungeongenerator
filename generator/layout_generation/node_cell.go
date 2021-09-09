package layout_generation

type nodeCell struct {
	nodeName string
	nodeTags []string // Locked? Key-containing? Anything.
}

func (n *nodeCell) setTags(tags []string) { // TODO: make the tags an array
	n.nodeTags = tags
}
