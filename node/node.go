package node

import (
  "osm-graph/tag"

  "github.com/paulmach/osm"
)

type Node struct {
  ID   int
  Lat  float64
  Lng  float64
  Tags tag.Tag
}

type NodesMap map[int]*Node

type Nodes []Node

func FromOSMNode(n osm.Node) Node {
  return Node{
    ID:   int(n.ID),
    Lat:  n.Lat,
    Lng:  n.Lon,
    Tags: tag.FromOSMTags(n.Tags),
  }
}

func FromOSMRelation(
  nn NodesMap,
  r osm.Relation,
  rr map[int64]osm.Relation,
  ways map[int64]osm.Way,
  aux Nodes) Nodes {
  for i := 0; i < len(r.Members); i++ {
    //TODO: make this as a go routine.
    m := r.Members[i]
    if m.Type == "node" {
      if v, ok := nn[int(m.Ref)]; ok {
        aux = append(aux, *v)
      }
    }
    if m.Type == "way" {
      w := ways[m.Ref]
      aux = append(aux, FromWay(w, nn)...)
    }
    if m.Type == "relation" {
      if v, ok := rr[m.Ref]; ok {
        aux = append(aux, FromOSMRelation(nn, v, rr, ways, Nodes{})...)
      }
    }
  }
  return aux
}

func FromWay(w osm.Way, nn NodesMap) Nodes {
  r := make(Nodes, 0)
  for i := range w.Nodes {
    if v, ok := nn[int(w.Nodes[i].ID)]; ok {
      r = append(r, *v)
    }
  }
  return r
}

func (nodes Nodes) IDs() []int {
  r := make([]int, 0)
  for i := range nodes {
    r = append(r, nodes[i].ID)
  }
  return r
}

//
//func (nodes NodesMap) ToGeojson() {
//  fc := geojson.NewFeatureCollection()
//  for _, n := range nodes {
//    fc.AddFeature(geojson.NewPointFeature([]float64{n.Lng, n.Lat}))
//  }
//  json.Write("nodes.json", fc)
//}
