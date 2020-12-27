package graph

import "fmt"

type AdjacencyListEdge struct {
	dst  int
	cost int
	next *AdjacencyListEdge
}

type AdjacencyListGraph struct {
	count int
	edges []*AdjacencyListEdge
}

func (a *AdjacencyListGraph) Init(count int) {
	a.count = count
	a.edges = make([]*AdjacencyListEdge, count)
}

func (a *AdjacencyListGraph) AddDirectedWeightedEdge(src, dst, cost int) {
	edge := &AdjacencyListEdge{dst, cost, a.edges[src]}
	a.edges[src] = edge
}

func (a *AdjacencyListGraph) AddUndirectedWeightedEdge(src, dst, cost int) {
	a.AddDirectedWeightedEdge(src, dst, cost)
	a.AddDirectedWeightedEdge(dst, src, cost)
}

func (a *AdjacencyListGraph) AddDirectedUnweightedEdge(src, dst int) {
	a.AddDirectedWeightedEdge(src, dst, 1)
}

func (a *AdjacencyListGraph) AddUndirectedUnweightedEdge(src, dst int) {
	a.AddUndirectedWeightedEdge(src, dst, 1)
}

func (a *AdjacencyListGraph) Print() {
	for i := 0; i < a.count; i++ {
		node := a.edges[i]
		for node != nil {
			fmt.Printf("Node %v, is connected to %v", i, node.dst)
			node = node.next
		}
		fmt.Println("")
	}
}
