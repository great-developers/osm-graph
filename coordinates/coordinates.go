package coordinates

import "github.com/umahmood/haversine"

type Coordinates struct {
  Lat, Lng float64
}

func Distance(a, b Coordinates) float64 {
  _, km := haversine.Distance(
    haversine.Coord{Lat: a.Lat, Lon: a.Lng},
    haversine.Coord{Lat: b.Lat, Lon: b.Lng},
  )
  return km * 1000
}
