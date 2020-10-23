package graph

import (
  jsonEncode "encoding/json"
  "io/ioutil"
  "log"
  "os"

  "github.com/JesseleDuran/osm-graph/file/json"
  "github.com/JesseleDuran/osm-graph/graph/edge"
  "github.com/JesseleDuran/osm-graph/graph/node"
  "github.com/golang/geo/s2"
)

type Graph struct {
  Nodes node.Nodes
}

func FromJSONGraphFile(path string) Graph {
  g := Graph{
    Nodes: make(node.Nodes, 0),
  }
  graphFile := struct {
    Edges [][]uint64 `json:"edges"`
  }{}
  file, err := ioutil.ReadFile(path)
  if err != nil {
    log.Printf("Couldn't read dump: %s", err.Error())
  }
  err = json.Read(&graphFile, file)
  for _, e := range graphFile.Edges {
    for i := 0; i < len(e)-1; i++ {
      source := node.Node{
        ID:       s2.CellID(e[i]),
        Neighbor: nil,
      }
      destiny := node.Node{
        ID:       s2.CellID(e[i+1]),
        Neighbor: nil,
      }

      source.Neighbor[destiny.ID] = edge.Edge{
        Weight: 2,
      }

      destiny.Neighbor[source.ID] = edge.Edge{
        Weight: 2,
      }

      g.Nodes[source.ID] = source
      g.Nodes[destiny.ID] = destiny
    }
  }
  return g
}

func FromJSONGraphFileStream(path string) Graph {
  g := Graph{
    Nodes: make(node.Nodes, 0),
  }
  jsonFile, _ := os.Open(path)
  defer jsonFile.Close()
  dec := jsonEncode.NewDecoder(jsonFile)
  _, err := dec.Token()

  if err != nil {
    log.Printf("Error reading open bracket [ during JSON parsing %s", err.Error())
    return g
  }

  for dec.More() {
    e := make([]uint64, 0)
    err := dec.Decode(&e)
    if err != nil {
      log.Printf("Error reading JSON structure during parsing %s", err.Error())
      continue
      //return g
    }
    for i := 0; i < len(e)-1; i++ {

      source := node.Node{
        ID:       s2.CellID(e[i]),
        Neighbor: make(edge.Edges, 0),
      }
      destiny := node.Node{
        ID:       s2.CellID(e[i+1]),
        Neighbor: make(edge.Edges, 0),
      }

      source.Neighbor[destiny.ID] = edge.Edge{
        Weight: 2,
      }

      destiny.Neighbor[source.ID] = edge.Edge{
        Weight: 2,
      }

      g.Nodes[source.ID] = source
      g.Nodes[destiny.ID] = destiny
    }
  }

  return g
}
