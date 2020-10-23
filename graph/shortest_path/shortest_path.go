package shortest_path

import (
  "github.com/JesseleDuran/osm-graph/coordinates"
)

type ShortestPath interface {
  ShortestPath(start, end coordinates.Coordinates)
}
