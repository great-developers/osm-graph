package json

import (
  "encoding/json"
  "fmt"
  "log"
  "os"
)

func Write(name string, content interface{}) string {
  f, err := os.Create(name)
  if err != nil {
    fmt.Println(err)
    return ""
  }
  d2, _ := json.Marshal(content)
  n2, err := f.Write(d2)
  if err != nil {
    log.Println(err)
    f.Close()
    return ""
  }
  log.Println(n2, "bytes written successfully")
  err = f.Close()
  if err != nil {
    log.Println(err)
    return ""
  }
  return f.Name()
}
