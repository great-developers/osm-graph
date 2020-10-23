package dijkstra

import (
  "log"
  "testing"
  "time"

  "github.com/JesseleDuran/osm-graph/graph"
)

func TestDijkstra(t *testing.T) {
  start := time.Now()
  g := graph.FromJSONGraphFileStream("graph/testdata/osm-graph-1.json")
  end := time.Since(start)
  log.Println("done graph", end.Milliseconds(), len(g.Nodes))
  time.Sleep(2*time.Hour)
  //weight, prev := Dijkstra(coordinates.Coordinates{
  //  Lat: 6.2451111,
  //  Lng: -75.5907088,
  //}, coordinates.Coordinates{
  //  Lat: 6.2451111,
  //  Lng: -75.5907088,
  //}, g)
  //
  //fmt.Println("----------------------------")
  //r := make([]int, 0)
  //for k, v := range weight {
  //  if v != math.MaxInt64 {
  //    //fmt.Println("k", k, v)
  //    //fmt.Println("rev", prev[k])
  //    r = append(r, prev[k])
  //  }
  //}
  //sort.Ints(r)
  //fmt.Println(r)
  //fmt.Println("----------------------------")
  //
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
