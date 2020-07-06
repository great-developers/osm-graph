package osm

import (
  "context"
  "os"
  "osm-graph/graph"
  "osm-graph/node"

  "github.com/paulmach/osm"
  "github.com/paulmach/osm/osmxml"
)

func FromFile(path string) graph.Graph {
  g := graph.Graph{}
  f, err := os.Open(path)
  if err != nil {
    panic(err)
  }
  defer f.Close()

  scanner := osmxml.New(context.Background(), f)
  defer scanner.Close()

  for scanner.Scan() {
    o := scanner.Object()
    if o.ObjectID().Type() == "node" {
      g.AddNode(node.FromOSMNode(*o.(*osm.Node)))
    }
    if o.ObjectID().Type() == "way" {
      nodes := o.(*osm.Way).Nodes
      for i := 0; i < len(nodes)-1; i++ {
        g.AddEdge(int64(nodes[i].ID), int64(nodes[i+1].ID))
      }
    }
  }

  scanErr := scanner.Err()
  if scanErr != nil {
    panic(scanErr)
  }
  return g
}
