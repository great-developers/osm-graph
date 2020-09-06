package graph

import (
  "context"
  "os"

  "github.com/JesseleDuran/osm-graph/edge"
  "github.com/JesseleDuran/osm-graph/node"
  "github.com/JesseleDuran/osm-graph/resources"
  "github.com/JesseleDuran/osm-graph/tag"
  "github.com/paulmach/osm"
  "github.com/paulmach/osm/osmxml"
)

//Graph it's the representation to be able to travel a geographical space.
type Graph struct {
  Nodes node.NodesMap
  Edges edge.Edges
}

func FromOSMFile(path string, weight edge.Weight) (Graph, error) {
  g := Graph{}
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
      auxTags := tag.FromOSMTags(r.Tags)
      if _, ok := auxTags["building"]; !ok {
        resources.Relations[r.ID.FeatureID().Ref()] = r
      }

    case "way":
      w := *o.(*osm.Way)
      resources.Ways[w.ID.FeatureID().Ref()] = w

    default:
      continue
    }
  }
  scanErr := scanner.Err()
  if scanErr != nil {
    return g, scanErr
  }

  var auxE []edge.Edges
  for _, v := range resources.Relations {
    auxE = append(auxE, edge.FromOSMRelation(v, g.Nodes, weight))
  }
  for _, v := range resources.Ways {
    auxE = append(auxE, edge.FromWays(v, g.Nodes, weight))
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
