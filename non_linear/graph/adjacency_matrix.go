package graph

import "fmt"

type AdjacencyMatrixGraph struct {
	count int
	adj   [][]int
}

func (a *AdjacencyMatrixGraph) Init(count int) {
	a.count = count
	a.adj = make([][]int, count)
	for row := range a.adj {
		a.adj[row] = make([]int, count)
	}
}

func (a *AdjacencyMatrixGraph) AddDirectedEdge(src, dst, cost int) {
	a.adj[src][dst] = cost
}

func (a *AdjacencyMatrixGraph) AddUndirectedEdge(src, dst, cost int) {
	a.AddDirectedEdge(src, dst, cost)
	a.AddDirectedEdge(dst, src, cost)
}

func (a *AdjacencyMatrixGraph) Print() {
	for i := 0; i < a.count; i++ {
		for j := 0; j < a.count; j++ {
			if a.adj[i][j] != 0 {
				fmt.Printf("Node %v is connected to: %v\n", i, j)
			}
		}
	}
}
