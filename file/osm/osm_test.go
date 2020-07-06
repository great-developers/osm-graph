package osm

import "testing"

func TestFromFile(t *testing.T) {
  g := FromFile("./map-el-poblado.osm")
  g.String()
}
