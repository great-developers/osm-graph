package dijkstra

import (
  "math"

  "github.com/JesseleDuran/osm-graph/coordinates"
  "github.com/JesseleDuran/osm-graph/graph"
  "github.com/JesseleDuran/osm-graph/graph/shortest_path/dijkstra/heap"
  "github.com/golang/geo/s2"
)

const INFINITE = math.MaxInt64

type Dijkstra struct {
  graph graph.Graph
}

type Previous map[s2.CellID]s2.CellID
type PathWeight map[s2.CellID]float64

func (d Dijkstra) ShortestPath(start, end coordinates.Coordinates) (PathWeight, Previous) {
  startCellID, endCellID := start.ToToken(), end.ToToken()
  return d.FromTokens(startCellID, endCellID)
}

func (d Dijkstra) FromTokens(start, end s2.CellID) (PathWeight, Previous) {
  //maps from each node to the total weight of the total shortest path.
  pathWeight := make(map[s2.CellID]float64, 0)

  //maps from each node to the previous node in the "current" shortest path.
  previous := make(map[s2.CellID]s2.CellID, 0)

  remaining := heap.Create()
  // insert first node id the PQ, the start node.
  remaining.Insert(heap.Node{Value: start, Cost: 0})

  // initialize pathWeight all to infinite value.
  for _, v := range d.graph.Nodes {
    pathWeight[v.ID] = INFINITE
  }
  //start node distance to itself is 0.
  pathWeight[start] = 0

  //the previous node does not exists
  previous[start] = INFINITE

  visit := make(map[s2.CellID]bool, 0)

  //while the PQ is not empty.
  for !remaining.IsEmpty() {
    // extract the min value of the PQ.
    min, _ := remaining.Min()
    visit[min.Value] = true
    remaining.DeleteMin()
    if min.Value == end {
      break
    }
    // if the node has edged, the loop through it.
    if v, ok := d.graph.Nodes[min.Value]; ok {
      //change to normal for
      for nodeNeighbor, e := range v.Neighbors {

        if visit[nodeNeighbor] {
          continue //change to negative condition
        }

        // the value is the one of the current node plus the weight(a, neighbor)
        currentPathValue := pathWeight[min.Value] + e.Weight

        if currentPathValue < pathWeight[nodeNeighbor] {
          pathWeight[nodeNeighbor] = currentPathValue
          previous[nodeNeighbor] = min.Value
        }
        remaining.Insert(heap.Node{Value: nodeNeighbor, Cost: currentPathValue})
      }
    }
  }
  return pathWeight, previous
}

//key : end, value: prev
func Path(start, end int, previous map[int]int) []int {
  result := make([]int, 0)
  result = append(result, end)
  prev := 0
  for prev != start {
    prev = previous[end]
    result = append(result, prev)
    end = prev
  }
  return result
}

//func CoordinatesPath(start, end coordinates.Coordinates, previous map[int]int,
//  g graph.Graph) []coordinates.Coordinates {
//  result := make([]coordinates.Coordinates, 0)
//  s, e := start.ToToken(), end.ToToken()
//  ms, me := g.NodesToCellID[s], g.NodesToCellID[e]
//  result = append(result, g.Nodes[me].Point)
//  prev := 0
//  for prev != ms {
//    prev = previous[me]
//    result = append(result, g.Nodes[prev].Point)
//    me = prev
//  }
//  return result
//}
