package graph

import (
	"testing"
)

var g Graph

func fillGraph() {
	nA := GraphNode{"A"}
	nB := GraphNode{"B"}
	nC := GraphNode{"C"}
	nD := GraphNode{"D"}
	nE := GraphNode{"E"}
	nF := GraphNode{"F"}
	g.AddNode(&nA)
	g.AddNode(&nB)
	g.AddNode(&nC)
	g.AddNode(&nD)
	g.AddNode(&nE)
	g.AddNode(&nF)

	g.AddEdge(&nA, &nB)
	g.AddEdge(&nA, &nC)
	g.AddEdge(&nB, &nE)
	g.AddEdge(&nC, &nE)
	g.AddEdge(&nE, &nF)
	g.AddEdge(&nD, &nA)
}

func TestGraph(t *testing.T) {
	fillGraph()
	g.Print()
}
