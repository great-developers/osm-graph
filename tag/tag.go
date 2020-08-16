package tag

import "github.com/paulmach/osm"

type Tag map[string]string

func FromOSMTags(tt osm.Tags) Tag {
  tag := make(Tag)
  for _, t := range tt {
    tag[t.Key] = t.Value
  }
  return tag
}
