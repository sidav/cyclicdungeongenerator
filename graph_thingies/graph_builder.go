package graph_thingies

var levelgraph = Graph{}

func makeInitialGraph() {
	start := &Node{name: "start"}
	finish := &Node{name: "finish"}
	interconnectTwoNodes(start, finish)
	levelgraph.addNode(start)
	levelgraph.addNode(finish)
}

func insertNodeToConnection() {
	alteration := &Node{name: "addition"}
	levelgraph.addNode(alteration)
}
