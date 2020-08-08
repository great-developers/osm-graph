package node

import (
  "osm-graph/file/json"

  geojson "github.com/paulmach/go.geojson"
  "github.com/paulmach/osm"
)

type Node struct {
  ID   int64
  Lat  float64
  Lng  float64
  Tags Tags // proper info of the node.
}

type Nodes map[int64]*Node

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

func (nodes Nodes) ToGeojson() {
  fc := geojson.NewFeatureCollection()
  for _, n := range nodes {
    fc.AddFeature(geojson.NewPointFeature([]float64{n.Lng, n.Lat}))
  }
  json.Write("nodes.json", fc)
}
