package main

import (
  "os"
  "fmt"
)

func main () {
  if len(os.Args) < 2 {
    fmt.Println("")
  } else {
    for _, arg := range os.Args[1:] {
      fmt.Println(arg)
    }
  }
}
