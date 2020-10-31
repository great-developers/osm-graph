package shortest_path

import "github.com/JesseleDuran/osm-graph/coordinates"

type Algorithm interface {
  ShortestPath(start, end coordinates.Coordinates) Response
}
