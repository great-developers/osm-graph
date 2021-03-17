package graph

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/JesseleDuran/osm-graph/store"
	"github.com/golang/geo/s2"
)

func TestGraph_AddNodes(t *testing.T) {
	g := newEmptyGraph()

	var table = []struct {
		in  Node
		out int
	}{
		{Node{Edges: nil}, 1},
	}

	for _, tt := range table {
		g.AddNode(24, tt.in)
		if len(g.Nodes) != tt.out {
			t.Errorf("got %v, want %v", tt.in, tt.out)
		}
	}

}

func TestGraph_FindNodeOrCreate(t *testing.T) {
	g := newEmptyGraph()

	var table = []struct {
		in  s2.CellID
		out bool
	}{
		{s2.CellID(23), true},
	}

	for _, tt := range table {
		g.FindNodeOrCreate(tt.in)
		if _, ok := g.Nodes[tt.in]; ok && ok != tt.out {
			t.Errorf("got %v, want %v", tt.in, tt.out)
		}
	}

}

func TestGraph_RelateNodes(t *testing.T) {
	idA, idB := s2.CellIDFromToken("94ce50"), s2.CellIDFromToken("94ce50")
	g := newEmptyGraph()
	g.RelateNodesByID(idA, idB)

	nodeA, nodeB := g.FindNodeOrCreate(idA), g.FindNodeOrCreate(idB)

	if !nodeA.IsRelated(idB) || !nodeB.IsRelated(idA) {
		t.Error("The nodes are not related", nodeA.Edges, nodeB.Edges)
	}
}

func TestBuildFromJsonFile(t *testing.T) {
	g := BuildFromJsonFile("testdata/osm-graph-sp-17.json")

	stores := store.FromCSV("testdata/sp-stores.csv")
	log.Println(len(stores))
	count := 0
	tokens := make(map[string]bool)
	for _, s := range stores {
		if !g.AddStore(s) {
			ll := s2.LatLngFromDegrees(s.Lat, s.Lng)
			token := s2.CellIDFromLatLng(ll).Parent(17).ToToken()
			tokens[token] = true
			count++
		}
	}

	for k := range tokens {
		fmt.Println(k)
	}

	log.Println("the graph needs:", len(tokens))
	time.Sleep(time.Hour)
}
