package node

import "github.com/paulmach/osm"

type Node struct {
  ID   int64
  Lat  float64
  Lng  float64
  Tags Tags // proper info of the node.
}

type Tags map[string]string

func FromOSMNode(n osm.Node) Node {
  tags := make(Tags)
  for _, t := range n.Tags {
    tags[t.Key] = t.Value
  }
  return Node{
    ID:   int64(n.ID),
    Lat:  n.Lat,
    Lng:  n.Lon,
    Tags: tags,
  }
}
