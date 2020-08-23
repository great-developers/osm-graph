package edge

import (
  "osm-graph/node"
  "osm-graph/tag"

  "github.com/paulmach/osm"
  "github.com/umahmood/haversine"
)

type MeansOfTransport int

const (
  Car MeansOfTransport = iota //same for motorcycle
  Bus
  Foot
  Bicycle
)

type Edge struct {
  SourceID  int
  DestinyID int
  Weight    float64
  Transport map[MeansOfTransport]bool
  Tags      tag.Tag
}

// Edges represents many "Edge" where the key of the first map is the ID of
// the Source node. The key of the second map is the ID of the Destiny node,
// which gives the resulting Edge between both nodes as a value.
// In this way it is optimized to apply Dijkstra.
type Edges map[int]map[int]*Edge

func FromOSMRelation(
  r osm.Relation,
  rr map[int64]osm.Relation,
  ways map[int64]osm.Way,
  nn node.NodesMap) Edges {
  edges := make(Edges)
  nodes := node.FromOSMRelation(nn, r, rr, ways, []node.Node{})
  for i := 0; i < len(nodes)-1; i++ {
    source := nodes[i]
    destiny := nodes[i+1]

    if source.ID == destiny.ID {
      continue
    }

    _, km := haversine.Distance(
      haversine.Coord{Lat: source.Lat, Lon: source.Lng},
      haversine.Coord{Lat: destiny.Lat, Lon: destiny.Lng},
    )
    meters := km * 1000

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
      Weight:    meters,
      Transport: nil,
      Tags:      tag.FromOSMTags(r.Tags),
    }

    edges[destiny.ID][source.ID] = &Edge{
      SourceID:  destiny.ID,
      DestinyID: source.ID,
      Weight:    meters,
      Transport: nil,
      Tags:      tag.FromOSMTags(r.Tags),
    }

  }
  return edges
}

func FromWays(w osm.Way, nn node.NodesMap) Edges {
  edges := make(Edges)
  nodes := node.FromWay(w, nn)

  for i := 0; i < len(nodes)-1; i++ {
    source := nodes[i]
    destiny := nodes[i+1]

    _, km := haversine.Distance(
      haversine.Coord{Lat: source.Lat, Lon: source.Lng},
      haversine.Coord{Lat: destiny.Lat, Lon: destiny.Lng},
    )
    meters := km * 1000

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
      Weight:    meters,
      Transport: nil,
      Tags:      tag.FromOSMTags(w.Tags),
    }

    edges[destiny.ID][source.ID] = &Edge{
      SourceID:  destiny.ID,
      DestinyID: source.ID,
      Weight:    meters,
      Transport: nil,
      Tags:      tag.FromOSMTags(w.Tags),
    }
  }
  return edges
}
