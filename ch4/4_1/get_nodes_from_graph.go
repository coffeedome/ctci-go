package main

func (g *Graph) GetNodes() []Node {
	graphNodes := []Node{}

	for {
		graphNodes := append(graphNodes, g.start)
		if g.start.edges == nil {
			return graphNodes
		}
		//depth-first search
		for _, n := range g.start.edges {
			graphNodes = append(graphNodes, n)
			g.start = n
		}
	}

}
