package graph

import (
	"testing"
)

func TestAdjListGraph(t *testing.T) {
	g := new(AdjacencyListGraph)
	g.Init(4)
	g.AddUndirectedWeightedEdge(0, 1, 1)
	g.AddUndirectedWeightedEdge(0, 2, 1)
	g.AddUndirectedWeightedEdge(1, 2, 1)
	g.AddUndirectedWeightedEdge(2, 3, 1)
	g.Print()
}
