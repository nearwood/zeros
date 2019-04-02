package main

import (
  "log"
  "os"
  "io"
  "fmt"
)

func main() {
  if len(os.Args) < 2 {
    log.Fatal("Specify filename")
    return
  }

  filename := os.Args[1]

  f, err := os.Open(filename)
  if err != nil {
    log.Fatal(err)
  }

  count := 0
  total := 0
  data := make([]byte, 4096)

  for {
    n, err := f.Read(data);

    if n > 0 {
      total += n
      for _, b := range data {
        if b == 0x0 {
          count += 1
        }
      }
    }

    if err == io.EOF {
      break;
    }

    if err != nil {
      log.Fatal(err)
    }
  }

  //fmt.Printf("Zero byte count: %d\n", count)
  //fmt.Printf("Total file bytes: %d\n", total)
  fmt.Printf("%3.3f%% empty\n", float64(count)/float64(total) * 100)

  if err := f.Close(); err != nil {
    log.Fatal(err)
  }

}

