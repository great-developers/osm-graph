package shortest_path

import (
  "math"
  "osm-graph/graph"
  "osm-graph/shortest_path/heap"
  "osm-graph/shortest_path/node"
)

const INFINITE = math.MaxInt64

func Dijkstra(start, end int, g graph.Graph) (map[int]float64, map[int]int) {
  //maps from each node to the total weight of the total shortest path.
  pathWeight := make(map[int]float64, 0)

  //maps from each node to the previous node in the "current" shortest path.
  previous := make(map[int]int, 0)

  remaining := heap.CreateN(100)
  // insert first node id the PQ, the start node.
  remaining.Insert(node.Node{Value: start, Cost: 0})

  // initialize pathWeight all to infinite value.
  for _, v := range g.Nodes {
    pathWeight[v.ID] = INFINITE
  }
  //start node distance to itself is 0.
  pathWeight[start] = 0

  //the previous node does not exists
  previous[start] = INFINITE

  visit := make(map[int]bool, 0)

  //while the PQ is not empty.
  for !remaining.IsEmpty() {
    // extract the min value of the PQ.
    min, _ := remaining.Min()
    visit[min.Value] = true
    remaining.DeleteMin()
    // if the node has edged, the loop through it.
    if v, ok := g.Edges[min.Value]; ok {
      //change to normal for
      for _, e := range v {

        if visit[e.DestinyID] {
          continue //change to negative condition
        }

        // the value is the one of the current node plus the weight(a, neighbor)
        currentPathValue := pathWeight[min.Value] + e.Weight

        if currentPathValue < pathWeight[e.DestinyID] {
          pathWeight[e.DestinyID] = currentPathValue
          previous[e.DestinyID] = min.Value
        }
        remaining.Insert(node.Node{Value: e.DestinyID, Cost: currentPathValue})
      }
    }
    //log.Println("size:", remaining.size)

  }

  return pathWeight, previous
}
