package osm

import (
  "encoding/xml"
  "fmt"
  "io/ioutil"

  "github.com/paulmach/osm"
)

func Read(filename string, o *osm.OSM) error {
  data, err := ioutil.ReadFile(filename)
  if err != nil {
    return fmt.Errorf("unable to open file: %v", err)
  }
  err = xml.Unmarshal(data, o)
  if err != nil {
    return fmt.Errorf("unable to unmarshal data: %v", err)
  }
  return nil
}
