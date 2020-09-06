package node

import (
  "github.com/JesseleDuran/osm-graph/coordinates"
  "github.com/JesseleDuran/osm-graph/graph/aux"
  "github.com/JesseleDuran/osm-graph/tag"
  "github.com/paulmach/osm"
)

//Node represents a geographical point of interest in a map.
type Node struct {
  ID    int
  Point coordinates.Coordinates
  Tags  tag.Tag
}

//NodesMap is a map where the key is the ID of the Node,
//and the value a reference of it.
type NodesMap map[int]*Node

type Nodes []Node

func FromOSMNode(n osm.Node) Node {
  return Node{
    ID: int(n.ID),
    Point: coordinates.Coordinates{
      Lat: n.Lat,
      Lng: n.Lon,
    },
    Tags: tag.FromOSMTags(n.Tags),
  }
}

func FromOSMRelation(nn NodesMap, r osm.Relation, nodesAux Nodes) Nodes {
  for i := 0; i < len(r.Members); i++ {
    //TODO: make this as a go routine.
    m := r.Members[i]
    if m.Type == "node" {
      if v, ok := nn[int(m.Ref)]; ok {
        nodesAux = append(nodesAux, *v)
      }
    }
    if m.Type == "way" {
      w := aux.Ways[m.Ref]
      nodesAux = append(nodesAux, FromWay(w, nn)...)
    }
    if m.Type == "relation" {
      if v, ok := aux.Relations[m.Ref]; ok {
        nodesAux = append(nodesAux, FromOSMRelation(nn, v, Nodes{})...)
      }
    }
  }
  return nodesAux
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
