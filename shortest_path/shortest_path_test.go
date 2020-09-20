package shortest_path

import (
  "fmt"
  "log"
  "math"
  "sort"
  "testing"
  "time"

  "github.com/JesseleDuran/osm-graph/graph"
)

func TestDijkstra(t *testing.T) {
  start := time.Now()
  g, err := graph.FromOSMFile(
    "/Users/jesseleduran/Documents/secure route graph/osm-graph/graph"+
      "/testdata/map-el-poblado.osm", nil)
  if err != nil {
    t.Fatal(err)
  }
  end := time.Since(start)
  log.Println("done graph", end.Milliseconds(), len(g.Nodes))

  //start = time.Now()
  for i := 0; i <= 100; i++ {
    weight, prev := Dijkstra(3352122607, 0, g)
    //end = time.Since(start)
    //fmt.Println("done dijsk", end.Milliseconds())
    fmt.Println("----------------------------")
    r := make([]int, 0)
    for k, v := range weight {
      if v != math.MaxInt64 {
        //fmt.Println("k", k, v)
        //fmt.Println("rev", prev[k])
        r = append(r, prev[k])
      }
    }
    sort.Ints(r)
    fmt.Println(r)
    fmt.Println("----------------------------")
  }
}

func TestPath(t *testing.T) {
  start, end := 1, 5
  previous := map[int]int{
    1: 0,
    2: 1,
    3: 2,
    4: 3,
    5: 4,
  }
  path := Path(start, end, previous)
  log.Println(path)
}
