package shortest_path

import "github.com/JesseleDuran/osm-graph/coordinates"

type Response struct {
  Leg             []Leg
  TotalDistanceKM float64
}

type Leg struct {
  Points     [2]Point
  DistanceKM float64
}

type Point struct {
  Point coordinates.Coordinates
  Name  string
}
