package main

import (
  "fmt"
//  "math"
)

func main() {
  type ByteSize int64
  const (
    _ = iota // ignore first value by assigning to blank identifier
    KB ByteSize = 1<<(10*iota)
    MB
    GB
    TB
    PB
  )
  fmt.Println(KB, MB, GB, TB, PB)
  fmt.Printf("%10g %20g %20g %20g %20g\n",KB, MB, GB, TB, PB)
}
