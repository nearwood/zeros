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
  var pPercentThresholdArg = flag.Float64("threshold", 0.0, "output only if the file meets this threshold (0-100%) of zeros")
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
    os.Exit(1)
  }

  emptyFile := false
  count := 0
  total := 0
  data := make([]byte, 4096)

  for {
    n, err := f.Read(data);

    if n > 0 {
      total += n
      for i, b := range data {
        if i >= n {
          break
        }

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
      emptyFile = true
      break;
    }

    if err != nil {
      log.Fatal(err)
    }
  }

  var outputString = ""

  if (optPrintFilename) {
    outputString += fmt.Sprintf("%s: ", *pFilenameArg)
  }

  var percentZero float64
  
  if emptyFile {
    percentZero = 100
  } else {
    percentZero = float64(count)/float64(total) * 100
  }
  if percentZero >= optThreshold {
    outputString += fmt.Sprintf("%3.3f%% empty\n", percentZero)

    if (optByteCounts) {
      outputString += fmt.Sprintf("%d/%d bytes\n", count, total)
    }

    fmt.Print(outputString)
  }

  if err := f.Close(); err != nil {
    log.Fatal(err)
  }
}

