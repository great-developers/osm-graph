package edge

import (
  "github.com/JesseleDuran/osm-graph/coordinates"
  "github.com/JesseleDuran/osm-graph/node"
  "github.com/JesseleDuran/osm-graph/tag"
  "github.com/JesseleDuran/osm-graph/transport"
  "github.com/paulmach/osm"
)

//Edge represents the way nodes are going to be related.
type Edge struct {
  SourceID  int
  DestinyID int
  Weight    float64
  Transport map[transport.Mode]bool
  Tags      tag.Tag
}

type Weight func(coordinates.Coordinates, coordinates.Coordinates) float64

// Edges represents many "Edges" where the key of the first map is SourceID.
//The key of the second map is the DestinyID, which gives the resulting Edge
//between both nodes as a value. In this way it is optimized to apply Dijkstra.
type Edges map[int]map[int]*Edge

func FromOSMRelation(r osm.Relation, nn node.NodesMap, weight Weight) Edges {
  edges := make(Edges)
  nodes := node.FromOSMRelation(nn, r, []node.Node{})
  for i := 0; i < len(nodes)-1; i++ {
    source := nodes[i]
    destiny := nodes[i+1]
    if source.ID == destiny.ID {
      continue
    }
    w := 0.0
    if weight == nil {
      w = coordinates.Distance(source.Point, destiny.Point)
    } else {
      w = weight(source.Point, destiny.Point)
    }

    // checks if the source already has some edges, if not, initialize it.
    if _, ok := edges[source.ID]; !ok {
      edges[source.ID] = make(map[int]*Edge, 0)
    }

    // checks if the destiny already has some edges, if not, initialize it.
    if _, ok := edges[destiny.ID]; !ok {
      edges[destiny.ID] = make(map[int]*Edge, 0)
    }

    edges[source.ID][destiny.ID] = &Edge{
      SourceID:  source.ID,
      DestinyID: destiny.ID,
      Weight:    w,
      Transport: nil,
      Tags:      tag.FromOSMTags(r.Tags),
    }

    edges[destiny.ID][source.ID] = &Edge{
      SourceID:  destiny.ID,
      DestinyID: source.ID,
      Weight:    w,
      Transport: nil,
      Tags:      tag.FromOSMTags(r.Tags),
    }
  }
  return edges
}

func FromWays(way osm.Way, nn node.NodesMap, weight Weight) Edges {
  edges := make(Edges)
  nodes := node.FromWay(way, nn)

  for i := 0; i < len(nodes)-1; i++ {
    source := nodes[i]
    destiny := nodes[i+1]

    w := 0.0
    if weight == nil {
      w = coordinates.Distance(source.Point, destiny.Point)
    } else {
      w = weight(source.Point, destiny.Point)
    }

    // checks if the source already has some edges, if not, initialize it.
    if _, ok := edges[source.ID]; !ok {
      edges[source.ID] = make(map[int]*Edge, 0)
    }

    // checks if the destiny already has some edges, if not, initialize it.
    if _, ok := edges[destiny.ID]; !ok {
      edges[destiny.ID] = make(map[int]*Edge, 0)
    }

    edges[source.ID][destiny.ID] = &Edge{
      SourceID:  source.ID,
      DestinyID: destiny.ID,
      Weight:    w,
      Transport: nil,
      Tags:      tag.FromOSMTags(way.Tags),
    }

    edges[destiny.ID][source.ID] = &Edge{
      SourceID:  destiny.ID,
      DestinyID: source.ID,
      Weight:    w,
      Transport: nil,
      Tags:      tag.FromOSMTags(way.Tags),
    }
  }
  return edges
}
