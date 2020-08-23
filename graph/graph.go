package graph

import (
  "context"
  "os"

  "github.com/JesseleDuran/osm-graph/edge"
  "github.com/JesseleDuran/osm-graph/node"
  "github.com/paulmach/osm"
  "github.com/paulmach/osm/osmxml"
)

//Graph it's the representation to be able to travel a geographical space.
type Graph struct {
  Nodes node.NodesMap
  Edges edge.Edges
}

func FromOSMFile(path string) (Graph, error) {
  g := Graph{}
  ww := map[int64]osm.Way{}
  rr := map[int64]osm.Relation{}
  f, err := os.Open(path)
  if err != nil {
    return g, err
  }
  defer f.Close()

  scanner := osmxml.New(context.Background(), f)
  defer scanner.Close()

  for scanner.Scan() {
    o := scanner.Object()
    switch o.ObjectID().Type() {
    case "node":
      g.AddNode(node.FromOSMNode(*o.(*osm.Node)))

    case "relation":
      r := *o.(*osm.Relation)
      auxTags := map[string]bool{}
      for i := range r.Tags {
        auxTags[r.Tags[i].Key] = false
      }
      if _, ok := auxTags["building"]; !ok {
        rr[r.ID.FeatureID().Ref()] = r
      }

    case "way":
      w := *o.(*osm.Way)
      ww[w.ID.FeatureID().Ref()] = w

    default:
      continue
    }
  }
  scanErr := scanner.Err()
  if scanErr != nil {
    return g, scanErr
  }
  var auxE []edge.Edges
  for _, v := range rr {
    auxE = append(auxE, edge.FromOSMRelation(v, rr, ww, g.Nodes))
  }

  for _, v := range ww {
    auxE = append(auxE, edge.FromWays(v, g.Nodes))
  }

  //array to array
  g.AddEdges(auxE)
  return g, nil
}

func (g *Graph) AddEdges(edges []edge.Edges) {
  if g.Edges == nil {
    g.Edges = make(edge.Edges)
  }
  for i := range edges {
    for source, m := range edges[i] {

      if _, ok := g.Edges[source]; !ok {
        g.Edges[source] = make(map[int]*edge.Edge, 0)
      }

      for related, e := range m {

        if _, ok := g.Edges[related]; !ok {
          g.Edges[related] = make(map[int]*edge.Edge, 0)
        }

        if v, ok := g.Edges[source][related]; ok {
          if v.Weight > e.Weight {
            g.Edges[source][related] = e
          }
        } else {
          g.Edges[source][related] = e
        }

        if v, ok := g.Edges[related][source]; ok {
          if v.Weight > e.Weight {
            g.Edges[related][source] = e
          }
        } else {
          g.Edges[related][source] = e
        }
      }
    }
  }
}

// AddNode adds a node to the graph.
func (g *Graph) AddNode(n node.Node) {
  if g.Nodes == nil {
    g.Nodes = make(node.NodesMap)
  }
  g.Nodes[n.ID] = &n
}

func (g *Graph) String() {
  //s := ""
  //for k, _ := range g.NodesMap {
  //  s += fmt.Sprintf("%d  ->", k)
  //  near := g.Edges[k]
  //  for j := 0; j < len(near); j++ {
  //    s += fmt.Sprintf(" %d ", near[j])
  //  }
  //  s += "\n"
  //}
  //fmt.Println(s)
}
