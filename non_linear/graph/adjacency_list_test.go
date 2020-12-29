package graph

import (
	"reflect"
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
	g2 := new(AdjacencyListGraph)
	g2.Init(6)
	g2.AddDirectedUnweightedEdge(5, 2)
	g2.AddDirectedUnweightedEdge(5, 0)
	g2.AddDirectedUnweightedEdge(4, 0)
	g2.AddDirectedUnweightedEdge(4, 1)
	g2.AddDirectedUnweightedEdge(2, 3)
	g2.AddDirectedUnweightedEdge(3, 1)
	exp3 := []int{5, 4, 2, 3, 1, 0}
	res3 := g2.TopologicalSort()
	if !reflect.DeepEqual(res3, exp3) {
		t.Errorf("Expected %v, got %v", exp3, res3)
	}

}
