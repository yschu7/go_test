// How to use slice
package main

import "fmt"

func main() {
  arr := []int32{1,2,3,4,5}
  x := arr[:]
  fmt.Println(arr)
  fmt.Println(x)
  x[0] = 0
  fmt.Println(arr)
  fmt.Println(x)
  y := make([]int32,6)
  copy(y, arr)
  fmt.Println(y)
  x[3] = 9
  fmt.Println(arr)
  fmt.Println(x)
  fmt.Println(y)
  //z := make([]int32,10)
  //var z [10]int32
  var z []int32
  println(len(z))
  z = append(y,9,8,7,6,5,4,3,2,1)
  fmt.Println(z)
  println(len(z))
  var kk []int32
  kk = z
  println(kk)
  fmt.Println(kk)
  println(&kk)
  fmt.Println(&kk)
  kk = append(kk,888)
  fmt.Println(kk)
}
