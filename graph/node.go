package graph

import (
  "github.com/golang/geo/s2"
)

type Node struct {
  ID        s2.CellID
  Neighbors Edges
}

type Nodes map[s2.CellID]Node
