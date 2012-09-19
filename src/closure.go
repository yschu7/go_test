package main

import (
  "fmt"
  "math"
)

func main() {
  // function is value, too
  sum := func(a,b int) int {return a+b} (3, 4)
  sum += func(a,b int) int {return a*b} (3, 4)
  fmt.Println(sum) // 19

  hypot := func(x,y float64) float64 {
    return math.Sqrt(x*x + y*y)
  }
  fmt.Println(hypot(3,4)) // 5

  pos, neg := adder(), adder()
  for i := 0; i < 10; i++ {
    fmt.Println(pos(i), neg(-2*i))
  }

}

// all functions are full closures.
func adder() func(int) int {
  sum := 0
  return func(x int) int {
    sum += x
    return sum
  }
}

