package graph

import (
	"container/list"
	"encoding/gob"
	"log"
	"os"

	"github.com/JesseleDuran/osm-graph/coordinates"
	"github.com/JesseleDuran/osm-graph/json"
	"github.com/golang/geo/s2"
)

// Represents a set of nodes related between them.
type Graph struct {
	Nodes Nodes
}

// newEmptyGraph creates a graph with 0 vertices.
func newEmptyGraph() Graph {
	return Graph{Nodes: make(Nodes)}
}

// Build creates a new graph from a .gob file.
func BuildFromGob(path string) Graph {
	file, _ := os.Open(path)
	defer file.Close()
	decoder := gob.NewDecoder(file)
	var g Graph
	err := decoder.Decode(&g)
	if err != nil {
		log.Panic("Error reading file")
	}
	return g
}

// BuildFromJsonFile creates a new graph from a json file.
func BuildFromJsonFile(path string) Graph {
	g := newEmptyGraph()
	decoder, file := json.NewDecoder(path)
	defer file.Close()
	for decoder.More() {
		nodes := make([]EncodedNode, 0)
		if decoder.Decode(&nodes) == nil {
			for i := 0; i < len(nodes)-1; i++ {
				g.RelateNodesByID(
					s2.CellID(nodes[i].CellId), s2.CellID(nodes[i+1].CellId),
				)
			}
		}
	}
	decoder.Token()
	return g
}

// Add note to the given graph.
func (g *Graph) AddNode(id s2.CellID, node Node) {
	g.Nodes[id] = node
}

// FindOrCreateNode find a node on the graph by the given ID.
// if the node does not exists then is created.
func (g *Graph) FindNodeOrCreate(id s2.CellID) *Node {
	node, ok := g.Nodes[id]
	if !ok {
		node = Node{Edges: make(map[s2.CellID]Edge)}
		g.AddNode(id, node)
	}
	return &node
}

// RelateNodesByID relate two nodes using its IDs.
func (g *Graph) RelateNodesByID(a, b s2.CellID) {
	nodeA, nodeB := g.FindNodeOrCreate(a), g.FindNodeOrCreate(b)
	pointA := coordinates.FromS2LatLng(a.LatLng())
	pointB := coordinates.FromS2LatLng(b.LatLng())
	w := coordinates.Distance(pointA, pointB)
	edge := Edge{Weight: w}

	// The relation of nodes is bi-directional.
	if edgeA, ok := nodeA.Edges[b]; !ok || edgeA.Weight > w {
		nodeA.Edges[b] = edge
	}
	if edgeB, ok := nodeB.Edges[a]; !ok || edgeB.Weight > w {
		nodeB.Edges[a] = edge
	}
}

func (g Graph) BFS(start s2.CellID, m float64) Nodes {
	result := make(Nodes, 0)
	visited := make(map[s2.CellID]float64)
	queue := list.New()
	queue.PushBack(start)
	visited[start] = 0

	for queue.Len() > 0 {
		qnode := queue.Front()
		queue.Remove(qnode)
		cellID := qnode.Value.(s2.CellID)
		for k, e := range g.Nodes[cellID].Edges {
			if _, ok := visited[k]; !ok {
				currentWeight := visited[cellID] + e.Weight
				if currentWeight < m {
					visited[k] += currentWeight
					queue.PushBack(k)
					result[k] = g.Nodes[k]
				}
			}
		}
	}
	return result
}

// Encode a graph.
func (g Graph) Encode() {
	file, _ := os.Create("graph-sp-17.gob")
	defer file.Close()
	encoder := gob.NewEncoder(file)
	encoder.Encode(g)
}
