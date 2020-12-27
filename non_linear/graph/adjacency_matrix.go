package graph

import "fmt"

/*
Pros:
1. Queries to see if there is an edge between u and v is O(1) time
2. Removing an edge takes O(1) time
Cons:
1. Space is quadratic
2. Adding a verted takes O(n^2) time
*/

type AdjacencyMatrixGraph struct {
	count int
	adj   [][]int
}

// O(n) time, O(n^2) space complexity
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
