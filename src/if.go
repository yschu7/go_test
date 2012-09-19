package main

import (
  "fmt"
  "math"
  "math/cmplx"
)

func pow(x, n, lim float64) float64 {
  if v := math.Pow(x,n); v < lim {
    return v
  } else {
    fmt.Printf("%g >= %g\n", v, lim)
  }
  // can't use v here
  return lim
}

func main() {
  var z complex128 = cmplx.Sqrt(-5+12i)
  fmt.Println(pow(3,2,10))
  fmt.Println(pow(3,3,20))
  fmt.Printf("%T(%v)\n", z, z)
}

