package dijkstra

import (
  "log"
  "math"

  "github.com/JesseleDuran/osm-graph/coordinates"
  "github.com/JesseleDuran/osm-graph/graph"
  "github.com/JesseleDuran/osm-graph/graph/shortest_path"
  "github.com/JesseleDuran/osm-graph/graph/shortest_path/dijkstra/heap"
  "github.com/golang/geo/s2"
)

const INFINITE = math.MaxInt64

type Dijkstra struct {
  graph graph.Graph
}

type Previous map[s2.CellID]s2.CellID
type PathWeight map[s2.CellID]float64

func (d Dijkstra) ShortestPath(start, end coordinates.Coordinates) shortest_path.Response {
  startCellID, endCellID := start.ToToken(), end.ToToken()
  pw, p := d.FromCellIDs(startCellID, endCellID)
  return shortest_path.Response{
    Leg:         toLegs(startCellID, endCellID, p),
    TotalWeight: pw[endCellID],
  }
}

func (d Dijkstra) FromCellIDs(start, end s2.CellID) (PathWeight, Previous) {
  //maps from each node to the total weight of the total shortest path.
  pathWeight := make(PathWeight, 0)

  //maps from each node to the previous node in the "current" shortest path.
  previous := make(Previous, 0)

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
func path(start, end s2.CellID, previous Previous) []s2.CellID {
  result := make([]s2.CellID, 0)
  result = append(result, end)
  var prev s2.CellID
  _, startOk := previous[start]
  _, endOk := previous[end]
  if !startOk && !endOk {
    return result
  }

  for prev != start {
    prev = previous[end]
    result = append(result, prev)
    end = prev
    log.Println(prev, end)
  }

  resultSorted := make([]s2.CellID, len(result))
  j := 0
  for i := len(result) - 1; i >= 0; i-- {
    resultSorted[j] = result[i]
    j++
  }
  return resultSorted
}

func toLegs(start, end s2.CellID, previous Previous) shortest_path.Legs {
  path := path(start, end, previous)
  legs := make(shortest_path.Legs, len(path))
  for i := 0; i < len(path)-1; i++ {
    legs[i] = shortest_path.Leg{
      Points: [2]shortest_path.Point{
        {
          Point: coordinates.FromS2LatLng(path[i].LatLng()),
          Name:  "",
        },
        {
          Point: coordinates.FromS2LatLng(path[i+1].LatLng()),
          Name:  "",
        },
      },
    }
  }
  return legs
}
