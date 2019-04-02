package main

import (
  "log"
  "os"
)

func main() {
  f, err := os.Open("test.mp4")
  if err != nil {
    log.Fatal(err)
  }

  if err := f.Close(); err != nil {
    log.Fatal(err)
  }

}

