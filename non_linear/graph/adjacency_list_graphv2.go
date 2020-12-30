package graph

import (
	"fmt"
	"sync"
)

/*
https://flaviocopes.com/golang-data-structure-graph/
*/

type GraphNode struct {
	Val interface{}
}

func (n *GraphNode) String() string {
	return fmt.Sprintf("%v", n.Val)
}

type Graph struct {
	Count int
	Nodes []*GraphNode
	Edges map[GraphNode][]*GraphNode // Can be map instead of slice for node keys that are string type
	lock  sync.RWMutex
}

func (g *Graph) AddNode(node *GraphNode) {
	g.lock.Lock()
	g.Count += 1
	g.Nodes = append(g.Nodes, node)
	g.lock.Unlock()
}

func (g *Graph) AddEdge(src, dst *GraphNode) {
	g.lock.Lock()
	if g.Edges == nil {
		g.Edges = make(map[GraphNode][]*GraphNode)
	}
	g.Edges[*src] = append(g.Edges[*src], dst)
	g.Edges[*dst] = append(g.Edges[*dst], src)
	g.lock.Unlock()
}

func (g *Graph) Print() {
	g.lock.RLock()
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
