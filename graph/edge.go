package graph

import (
  "github.com/JesseleDuran/osm-graph/coordinates"
  "github.com/golang/geo/s2"
)

type Edge struct {
  Weight float64
}

type SetWeight func(coordinates.Coordinates, coordinates.Coordinates) float64

type Edges map[s2.CellID]Edge
