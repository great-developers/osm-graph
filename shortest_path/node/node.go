package node

//Node represents the item used in the shortest path with the minimum necessary
//information, the cost or weight and its ID or value.
type Node struct {
  Value int
  Cost  float64
}

type Nodes []Node
