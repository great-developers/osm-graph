package shortest_path

import (
  "github.com/JesseleDuran/osm-graph/coordinates"
)

type ShortestPath struct {
  Algorithm Algorithm
}

func FromAlgorithm(a Algorithm) *ShortestPath {
  return &ShortestPath{
    Algorithm: a,
  }
}

func (spa *ShortestPath) SetAlgorithm(a Algorithm) {
  spa.Algorithm = a
}

func (spa *ShortestPath) ShortestPath(start, end coordinates.Coordinates) Response {
  return spa.Algorithm.ShortestPath(start, end)
}
