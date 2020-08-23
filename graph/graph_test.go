package graph

import (
  "log"
  "osm-graph/node"
  "testing"
)

var g Graph

func fillGraph() {
  nA := node.Node{
    ID: 1,
  }
  nB := node.Node{
    ID: 2,
  }
  nC := node.Node{
    ID: 3,
  }
  nD := node.Node{
    ID: 4,
  }
  nE := node.Node{
    ID: 5,
  }
  nF := node.Node{
    ID: 6,
  }
  g.AddNode(nA)
  g.AddNode(nB)
  g.AddNode(nC)
  g.AddNode(nD)
  g.AddNode(nE)
  g.AddNode(nF)

  //g.AddEdge(1, 2)
  //g.AddEdge(1, 3)
  //g.AddEdge(2, 5)
  //g.AddEdge(3, 5)
  //g.AddEdge(5, 6)
  //g.AddEdge(4, 1)
}

func TestAdd(t *testing.T) {
  fillGraph()
  g.String()
}

func TestFromOSMFile(t *testing.T) {
  g, _ := FromOSMFile("testdata/Bogota.osm")
  //g.NodesMap.ToGeojson()
  log.Println("finished")
  log.Println(len(g.Nodes))
  log.Println(len(g.Edges[6069561818]))
}

//func TestFromOSMPBFFile(t *testing.T) {
//  FromOSMPBFFile("/Users/jesseleduran/Documents/secure route graph/osm-graph/graph/testdata/colombia-latest.osm.pbf")
//}
