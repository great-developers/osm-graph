package dijkstra

import (
  "fmt"
  "log"
  "math"
  "testing"
  "time"

  "github.com/JesseleDuran/osm-graph/graph"
  "github.com/golang/geo/s2"
)

func TestDijkstra(t *testing.T) {
  start := time.Now()
  g := graph.FromJSONGraphFileStream("testdata/osm-graph-1.json")
  end := time.Since(start)
  log.Println("done graph", end.Milliseconds(), len(g.Nodes))
  //for k, _ := range g.Nodes {
  // log.Println(k.ToToken())
  //}

  s := s2.CellIDFromToken("94ce595164")
  e := s2.CellIDFromToken("94ce50b26c")
  weight, prev := FromToken(s, e, g)
  //
  fmt.Println("----------------------------")
  r := make([]s2.CellID, 0)
  for k, v := range weight {
    if v != math.MaxInt64 {
      fmt.Println("k", k, v)
      fmt.Println("rev", prev[k])
      r = append(r, prev[k])
    }
  }
  fmt.Println(len(r))
  fmt.Println("----------------------------")

  //log.Println(Path(7126102609, r[0], prev))
  //c := g.Nodes[r[0]]

  //cc := CoordinatesPath(coordinates.Coordinates{
  //  Lat: 6.2451111,
  //  Lng: -75.5907088,
  //}, c.Point, prev, g)
  //for i := range cc {
  //  fmt.Println("[", cc[i].Lng, ",", cc[i].Lat, "],")
  //}
  //log.Println(weight[r[0]])

  //for i := 0; i <= 100; i++ {
  //  weight, prev := Dijkstra(7126102609, 0, g)
  //  fmt.Println("----------------------------")
  //  r := make([]int, 0)
  //  for k, v := range weight {
  //    if v != math.MaxInt64 {
  //      //fmt.Println("k", k, v)
  //      //fmt.Println("rev", prev[k])
  //      r = append(r, prev[k])
  //    }
  //  }
  //  sort.Ints(r)
  //  fmt.Println(r)
  //  fmt.Println("----------------------------")
  //}
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
