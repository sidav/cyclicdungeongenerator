package graph_thingies

type Graph struct {
	nodes []*Node
}

func (g *Graph) addNode(n *Node) {
	g.nodes = append(g.nodes, n)
}

func (g *Graph) getRandomNode() *Node {
	return nil
}
