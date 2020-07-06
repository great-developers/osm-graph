package graph

import (
  "osm-graph/node"
  "testing"
)

var g Graph

func fillGraph() {
  nA := node.Node{
    ID: 1,
  }
  nB := node.Node{
    ID: 2,
  }
  nC := node.Node{
    ID: 3,
  }
  nD := node.Node{
    ID: 4,
  }
  nE := node.Node{
    ID: 5,
  }
  nF := node.Node{
    ID: 6,
  }
  g.AddNode(nA)
  g.AddNode(nB)
  g.AddNode(nC)
  g.AddNode(nD)
  g.AddNode(nE)
  g.AddNode(nF)

  g.AddEdge(1, 2)
  g.AddEdge(1, 3)
  g.AddEdge(2, 5)
  g.AddEdge(3, 5)
  g.AddEdge(5, 6)
  g.AddEdge(4, 1)
}

func TestAdd(t *testing.T) {
  fillGraph()
  g.String()
}
