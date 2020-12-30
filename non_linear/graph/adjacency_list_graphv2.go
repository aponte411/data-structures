package graph

import (
	"fmt"
	que "github.com/aponte411/data-structures/linear/queue"
	stk "github.com/aponte411/data-structures/linear/stack"
	"sync"
)

/*
https://flaviocopes.com/golang-data-structure-graph/
*/

type GraphNode struct {
	val  interface{}
	cost int
}

type Graph struct {
	Count int
	Nodes []*GraphNode
	Edges map[GraphNode][]*GraphNode // Can be map instead of slice for node keys that are string type
	lock  sync.RWMutex
}

func (g *Graph) AddNode(node *GraphNode) {
	g.Count += 1
	g.lock.Lock()
	g.Nodes = append(g.Nodes, node)
	g.lock.Unlock()
}

func (g *Graph) AddEdge(src, dst *GraphNode) {
	g.lock.Lock()
	if g.Edges == nil {
		g.Edges = make(map[GraphNode][]*GraphNode)
	}
	g.Edges[*src] = append(g.Edges[*src], src)
	g.Edges[*dst] = append(g.Edges[*dst], dst)
	g.lock.Unlock()
}

func (g *Graph) Print() {
	g.lock.Rlock()
	s := ""
	for i := 0; i < g.Count; i++ {
		node := g.Nodes[i]
		s += node.String() + " -> "
		neighbors := g.Edges[*node]
		for j := 0; j < len(neighbors); j++ {
			s += neighbors[j].String() + " "
		}
		s += "\n"
	}
	fmt.Println(s)
	g.lock.RUnlock()
}
