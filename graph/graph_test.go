package graph

import (
  "log"
  "testing"
  "time"
)

func TestFromJSONGraphFile(t *testing.T) {
  g := FromJSONGraphFile("testdata/osm-graph.json")
  log.Println(len(g.Nodes))
  time.Sleep(3*time.Hour)
}

func TestFromJSONGraphFileStream(t *testing.T) {
  g := FromJSONGraphFileStream("testdata/osm-graph-1.json")
  log.Println(len(g.Nodes))
  time.Sleep(3*time.Hour)
}
