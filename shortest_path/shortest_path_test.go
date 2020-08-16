package shortest_path

import (
  "log"
  "math"
  "osm-graph/graph"
  "testing"
)

func TestDijkstra(t *testing.T) {
  g, err := graph.FromOSMFile("/Users/jesseleduran/Documents/secure route graph/osm-graph/graph/testdata/map-el-poblado.osm")
  if err != nil {
    t.Fatal(err)
  }

  weight, prev := Dijkstra(5532355787, 0, g)
  for k, v := range weight {
    if v != math.MaxInt64 {
      log.Println("----------------------------")
      log.Println("k", k, v)
      log.Println("rev", prev[k])
      log.Println("----------------------------")
    }
  }

}
