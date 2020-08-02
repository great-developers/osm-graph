// Package graph creates a Graph data structure for the Item type
package graph

import (
  "fmt"
  "osm-graph/node"
)

type Graph struct {
  nodes map[int64]*Node
  edges map[int64][]int64
}

// Node a single node that composes the graph.
type Node struct {
  Value node.Node
}

// AddNode adds a node to the graph.
func (g *Graph) AddNode(n node.Node) {
  if g.nodes == nil {
    g.nodes = make(map[int64]*Node)
  }
  g.nodes[n.ID] = &Node{Value: n}
}

// AddEdge adds an edge to the graph.
func (g *Graph) AddEdge(id1, id2 int64) {
  if g.edges == nil {
    g.edges = make(map[int64][]int64)
  }
  g.edges[id1] = append(g.edges[id1], id2)
  g.edges[id2] = append(g.edges[id2], id1)
}

func (g *Graph) String() {
  s := ""
  for k, _ := range g.nodes {
    s += fmt.Sprintf("%d  ->", k)
    near := g.edges[k]
    for j := 0; j < len(near); j++ {
      s += fmt.Sprintf(" %d ", near[j])
    }
    s += "\n"
  }
  fmt.Println(s)
}
