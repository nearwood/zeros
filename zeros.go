package main

import (
  "log"
  "os"
  "io"
  "fmt"
  "flag"
)

func main() {
  var pFilenameArg = flag.String("file", "", "file to parse")
  var pPrintFilenameArg = flag.Bool("print", true, "print filename in output")
  var pSkipNCArg = flag.Bool("skipnc", true, "skip non-contiguous bytes")
  var pPercentThresholdArg = flag.Float64("threshold", 0.0, "threshold (0-100%) to print results")
  var pByteCountsArg = flag.Bool("bytes", false, "print byte counts")
  
  flag.Parse()

  var optPrintFilename = *pPrintFilenameArg
  var optSkipNC = *pSkipNCArg
  var optThreshold = *pPercentThresholdArg
  var optByteCounts = *pByteCountsArg

  if optThreshold < 0.0 || optThreshold > 100.0 {
    log.Printf("Threshold must be 0-100, but was %3.3f\n", optThreshold)
    flag.PrintDefaults()
    os.Exit(1)
  }

  f, err := os.Open(*pFilenameArg)
  if err != nil {
    log.Print(err)
    flag.PrintDefaults()
    os.Exit(1)
  }

  count := 0
  total := 0
  data := make([]byte, 4096)

  for {
    n, err := f.Read(data);

    if n > 0 {
      total += n
      for i, b := range data {
        if i > 0 {
          if (optSkipNC) {
            if b == 0x0 && data[i - 1] == 0x0 {
              count += 1
            }
          } else {
            if b == 0x0 {
              count += 1
            }
          }
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
  var outputTemplate = "%3.3f%% empty\n"
  if (optPrintFilename) {
    outputTemplate = "%s: " + outputTemplate
  }

  var percentZero = float64(count)/float64(total) * 100
  if percentZero >= optThreshold {
    if (optByteCounts) {
      fmt.Printf(outputTemplate + "%d/%d bytes\n", *pFilenameArg, percentZero, count, total)
    } else {
      fmt.Printf(outputTemplate, *pFilenameArg, percentZero)
    }
  }

  if err := f.Close(); err != nil {
    log.Fatal(err)
  }

}

