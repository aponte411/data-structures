package graph

import (
	"testing"
)

func TestAdjMatrix(t *testing.T) {
	g := new(AdjacencyMatrixGraph)
	g.Init(4)
	g.AddUndirectedEdge(0, 1, 1)
	g.AddUndirectedEdge(0, 2, 1)
	g.AddUndirectedEdge(1, 2, 1)
	g.AddUndirectedEdge(2, 3, 1)
	g.Print()
}
