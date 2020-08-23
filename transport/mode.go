package transport

//Mode represents the means of transport by which the graph can filter
//its routes.
type Mode int

const (
  Car Mode = iota //same for motorcycle
  Foot
  Bicycle
)
