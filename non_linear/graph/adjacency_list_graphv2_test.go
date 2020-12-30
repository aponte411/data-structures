package graph

import (
	"testing"
)

var g Graph

func fillGraph() {
	nA := GraphNode{"David"}
	nB := GraphNode{"Pati"}
	nC := GraphNode{"Ryan"}
	nD := GraphNode{"Laura"}
	nE := GraphNode{"Peter"}
	nF := GraphNode{"Jason"}
	nG := GraphNode{"Jon"}
	nH := GraphNode{"Orlando"}
	nI := GraphNode{"Dagmar"}

	g.AddNode(&nA)
	g.AddNode(&nB)
	g.AddNode(&nC)
	g.AddNode(&nD)
	g.AddNode(&nE)
	g.AddNode(&nF)
	g.AddNode(&nG)
	g.AddNode(&nH)
	g.AddNode(&nI)

	g.AddEdge(&nA, &nB)
	g.AddEdge(&nA, &nC)
	g.AddEdge(&nB, &nE)
	g.AddEdge(&nC, &nE)
	g.AddEdge(&nE, &nF)
	g.AddEdge(&nD, &nA)
	g.AddEdge(&nH, &nI)
	g.AddEdge(&nH, &nF)
	g.AddEdge(&nI, &nB)
	g.AddEdge(&nG, &nF)

}

func TestGraph(t *testing.T) {
	fillGraph()
	g.Print()
}
