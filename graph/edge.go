package graph

import (
	"github.com/golang/geo/s2"
)

type Edge struct {
	Weight float64
}

type Edges map[s2.CellID]Edge
