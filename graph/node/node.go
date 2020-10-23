package node

import (
  "github.com/JesseleDuran/osm-graph/graph/edge"
  "github.com/golang/geo/s2"
)

type Node struct {
  ID       s2.CellID
  Neighbor edge.Edges
}

type Nodes map[s2.CellID]Node
