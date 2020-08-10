// Package graph creates a Graph data structure for the Item type
package graph

import (
  "context"
  "fmt"
  "os"
  "osm-graph/node"

  "github.com/paulmach/osm"
  "github.com/paulmach/osm/osmxml"
)

type Graph struct {
  Nodes node.Nodes
  Edges map[int64][]int64
}

func FromOSMFile(path string) Graph {
  g := Graph{}
  f, err := os.Open(path)
  if err != nil {
    panic(err)
  }
  defer f.Close()

  scanner := osmxml.New(context.Background(), f)
  defer scanner.Close()

  for scanner.Scan() {
    o := scanner.Object()
    if o.ObjectID().Type() == "node" {
      g.AddNode(node.FromOSMNode(*o.(*osm.Node)))
    }
    if o.ObjectID().Type() == "way" {
      nodes := o.(*osm.Way).Nodes
      for i := 0; i < len(nodes)-1; i++ {
        g.AddEdge(int64(nodes[i].ID), int64(nodes[i+1].ID))
      }
    }
  }

  scanErr := scanner.Err()
  if scanErr != nil {
    panic(scanErr)
  }
  return g
}

// AddNode adds a node to the graph.
func (g *Graph) AddNode(n node.Node) {
  if g.Nodes == nil {
    g.Nodes = make(node.Nodes)
  }
  g.Nodes[n.ID] = &n
}

// AddEdge adds an edge to the graph.
func (g *Graph) AddEdge(id1, id2 int64) {
  if g.Edges == nil {
    g.Edges = make(map[int64][]int64)
  }
  g.Edges[id1] = append(g.Edges[id1], id2)
  g.Edges[id2] = append(g.Edges[id2], id1)
}

func (g *Graph) String() {
  s := ""
  for k, _ := range g.Nodes {
    s += fmt.Sprintf("%d  ->", k)
    near := g.Edges[k]
    for j := 0; j < len(near); j++ {
      s += fmt.Sprintf(" %d ", near[j])
    }
    s += "\n"
  }
  fmt.Println(s)
}
