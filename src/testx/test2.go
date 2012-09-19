package main
import (
   "fmt"
)
var (
  a = 5
  b = 5.8
  c = false
  d = "hello"
  e = []int{1,2,3,45}
  變= "中文"
)

const (
  A = iota
  B
  C
  D
  E
)

func printVar() {
   fmt.Println(a,b,c,d,e)
   fmt.Println(變)
}
func printConst() {
  fmt.Println(A,B,C,D,E)
}

