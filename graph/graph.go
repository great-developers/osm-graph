package graph

import (
  "encoding/json"
  "log"
  "os"

  "github.com/golang/geo/s2"
)

type Graph struct {
  Nodes Nodes
}

func FromJSONGraphFileStream(path string) Graph {
  g := Graph{Nodes: make(Nodes, 0)}
  jsonFile, _ := os.Open(path)
  defer jsonFile.Close()
  dec := json.NewDecoder(jsonFile)
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
      source := Node{ID: s2.CellID(e[i]), Neighbors: make(Edges, 0)}
      destiny := Node{ID: s2.CellID(e[i+1]), Neighbors: make(Edges, 0)}

      source.Neighbors[destiny.ID] = Edge{Weight: 2}
      destiny.Neighbors[source.ID] = Edge{Weight: 2}

      g.Nodes[source.ID] = source
      g.Nodes[destiny.ID] = destiny
    }
  }
  return g
}
