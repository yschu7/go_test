package main
import (
  "fmt"
)
func zero(xPtr *int) {
    *xPtr = 0
}
func one(xPtr *int) {
    *xPtr = 1
}
func square(x *float64) {
    *x = *x * *x
}
func main() {
    x := 5
    zero(&x)
    fmt.Println(x) // x is 0
    xPtr := new(int) // allocate memory for int and return its addr
    one(xPtr)
    fmt.Println(*xPtr) // x is 1
    y := 1.5
    square(&y)
    fmt.Println(y)
}
