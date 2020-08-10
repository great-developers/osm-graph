package edge

import "osm-graph/node"

type MeansOfTransport int

const (
  Car MeansOfTransport = iota //same for motorcycle
  Bus
  Foot
  Bicycle
)

type Edge struct {
  Name      string
  Source    node.Node
  Destiny   node.Node
  Weight    int
  Transport MeansOfTransport
}

// Edges represents many "Edge" where the key of the first map is the ID of
// the Source node. The key of the second map is the ID of the Destiny node,
// which gives the resulting Edge between both nodes as a value.
// In this way it is optimized to apply Dijkstra.
type Edges map[int]map[int]*Edge
