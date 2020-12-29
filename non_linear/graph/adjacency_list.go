package graph

import (
	"fmt"
	que "github.com/aponte411/data-structures/linear/queue"
	stk "github.com/aponte411/data-structures/linear/stack"
)

type AdjacencyListEdge struct {
	dst  int
	cost int
	next *AdjacencyListEdge
}

type AdjacencyListGraph struct {
	count int
	edges []*AdjacencyListEdge // Can be map instead of slice for node keys that are string type
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
			fmt.Printf("Node %v, is connected to %v\n", i, node.dst)
			node = node.next
		}
	}
}

// DFS find bath between src and dst
// O(V + E) time, O(V + E)
func (a *AdjacencyListGraph) DFSFindPathBetween(src, dst int) bool {
	visited := make([]bool, a.count)
	a.DFS(src, visited)
	return visited[dst]
}

func (a *AdjacencyListGraph) DFS(src int, visited []bool) {
	visited[src] = true
	head := a.edges[src]
	for head != nil {
		if visited[head.dst] == false {
			a.DFS(head.dst, visited)
		}
		head = head.next
	}
}

func (a *AdjacencyListGraph) DFSStackFindPathBetween(src, dst int) bool {
	visited := make([]bool, a.count)
	stack := new(stk.StackLinkedList)
	visited[src] = true
	stack.Push(src)
	for !stack.IsEmpty() {
		curr := stack.Pop().(int)
		head := a.edges[curr]
		for head != nil {
			if visited[head.dst] == false {
				visited[head.dst] = true
				stack.Push(head.dst)
			}
			head = head.next
		}
	}
	return visited[dst]
}

func (a *AdjacencyListGraph) BFSFindPathBetween(src, dst int) bool {
	visited := make([]bool, a.count)
	queue := new(que.QueueLinkedList)
	visited[src] = true
	queue.Enqueue(src)
	for !queue.IsEmpty() {
		curr := queue.Dequeue().(int)
		head := a.edges[curr]
		for head != nil {
			if visited[head.dst] == false {
				visited[head.dst] = true
				queue.Enqueue(head.dst)
			}
			head = head.next
		}
	}
	return visited[dst]
}

// TODO: FIX
func (a *AdjacencyListGraph) TopologicalSort() {
	stack := new(stk.StackLinkedList)
	visited := make([]bool, a.count)
	for i := 0; i < a.count; i++ {
		if visited[i] == false {
			visited[i] = true
			a.TopologicalSortDFS(i, visited, stack)
		}
	}
	for !stack.IsEmpty() {
		fmt.Println(stack.Pop().(int))
	}
}

func (a *AdjacencyListGraph) TopologicalSortDFS(index int, visited []bool, indexStack stk.StackLinkedList) {
	head := a.edges[index]
	for head != nil {
		if visited[head.dst] == false {
			visited[head.dst] = true
			a.TopologicalSortDFS(head.dst, visited, indexStack)
		}
	}
	indexStack.Push(index)
}
