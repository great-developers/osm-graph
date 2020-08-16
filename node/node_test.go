package node

import (
  "log"
  "testing"

  "github.com/paulmach/osm"
)

func TestFromOSMRelation(t *testing.T) {
  r := osm.Relation{
    ID: 1,
    Members: osm.Members{
      {Type: "node", Ref: 12},
      {Type: "node", Ref: 22},
      {Type: "way", Ref: 13},
      {Type: "relation", Ref: 14},
    },
  }
  nn := NodesMap{
    12: &Node{ID: 12},
    22: &Node{ID: 22},
    5:  &Node{ID: 5},
    7:  &Node{ID: 7},
    1:  &Node{ID: 1},
    45: &Node{ID: 45}}

  rr := map[int64]osm.Relation{
    14: {ID: 14, Members: osm.Members{
      {Type: "node", Ref: 5},
      {Type: "way", Ref: 8},
    }},
  }
  ways := map[int64]osm.Way{
    13: {ID: 13, Nodes: osm.WayNodes{{ID: 7}}},
    8: {ID: 8, Nodes: osm.WayNodes{
      {ID: 1},
      {ID: 45},
    }},
  }
  result := FromOSMRelation(nn, r, rr, ways, Nodes{})
  log.Println(result)
}
