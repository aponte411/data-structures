package graph

import (
	"testing"
)

func TestAdjListGraph(t *testing.T) {
	g := new(AdjacencyListGraph)
	g.Init(8)
	g.AddUndirectedUnweightedEdge(0, 3)
	g.AddUndirectedUnweightedEdge(0, 2)
	g.AddUndirectedUnweightedEdge(0, 1)
	g.AddUndirectedUnweightedEdge(1, 4)
	g.AddUndirectedUnweightedEdge(2, 5)
	g.AddUndirectedUnweightedEdge(3, 6)
	g.AddUndirectedUnweightedEdge(6, 7)
	g.AddUndirectedUnweightedEdge(5, 7)
	g.AddUndirectedUnweightedEdge(4, 7)
	g.Print()
	res1 := g.DFSFindPathBetween(0, 6)
	res2 := g.BFSFindPathBetween(0, 6)
	if !res1 {
		t.Errorf("Expected path between node 0 and node 6, %v", res1)
	}
	if !res2 {
		t.Errorf("Expected path between node 0 and node 6, %v", res2)
	}
}