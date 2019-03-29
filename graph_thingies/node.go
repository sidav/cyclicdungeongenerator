package graph_thingies

const MAX_NEIGHBOURS = 4

type Node struct {
	name string
	connections []*Node // not more than 4 neighbours allowed, we're on grid-based map after all
}

func (n *Node) isNeighbourPlaceVacant() bool {
	return len(n.connections) < MAX_NEIGHBOURS
}

func interconnectTwoNodes(a, b *Node) {
	if a.isNeighbourPlaceVacant() && b.isNeighbourPlaceVacant() {
		a.connections = append(a.connections, b)
		b.connections = append(b.connections, a)
	}
}
