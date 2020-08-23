package shortest_path

import (
  "log"
  "math"
  "osm-graph/graph"
  "testing"
  "time"
)

func TestDijkstra(t *testing.T) {
  start := time.Now()
  g, err := graph.FromOSMFile("/Users/jesseleduran/Documents/secure route graph/osm-graph/graph/testdata/Bogota.osm")
  if err != nil {
    t.Fatal(err)
  }
  end := time.Since(start)
  log.Println("done graph", end.Milliseconds(), len(g.Nodes))

  start = time.Now()
  weight, prev := Dijkstra(5854938735, 0, g)
  end = time.Since(start)
  log.Println("done dijsk", end.Milliseconds())
  for k, v := range weight {
   if v != math.MaxInt64 {
     log.Println("----------------------------")
     log.Println("k", k, v)
     log.Println("rev", prev[k])
     log.Println("----------------------------")
   }
  }

}
