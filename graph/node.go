package graph

import (
	"github.com/JesseleDuran/osm-graph/json"
	"github.com/JesseleDuran/osm-graph/store"
	"github.com/golang/geo/s2"
	geojson "github.com/paulmach/go.geojson"
)

type Node struct {
	Edges  Edges
	Stores []store.Store
}

type EncodedNode struct {
	CellId uint64
}

type Nodes map[s2.CellID]Node

func (nodes Nodes) ToGeoJSON() {
	fc := geojson.NewFeatureCollection()
	for k, _ := range nodes {
		children := k.Children()
		f := geojson.NewPolygonFeature([][][]float64{
			{
				[]float64{children[0].LatLng().Lng.Degrees(), children[0].LatLng().Lat.Degrees()},
				[]float64{children[1].LatLng().Lng.Degrees(), children[1].LatLng().Lat.Degrees()},
				[]float64{children[2].LatLng().Lng.Degrees(), children[2].LatLng().Lat.Degrees()},
				[]float64{children[3].LatLng().Lng.Degrees(), children[3].LatLng().Lat.Degrees()},
				[]float64{children[0].LatLng().Lng.Degrees(), children[0].LatLng().Lat.Degrees()},
			},
		})
		f.ID = k
		f.Properties = map[string]interface{}{
			"id": k.ToToken(),
		}
		fc.AddFeature(f)
	}
	json.Write("nodes.json", fc)
}

func (n Node) IsRelated(id s2.CellID) bool {
	_, ok := n.Edges[id]
	return ok
}

func (n *Node) AddStore(s store.Store) {
	n.Stores = append(n.Stores, s)
}
