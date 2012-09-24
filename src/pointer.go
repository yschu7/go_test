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
    fmt.Println(y)  // 2.25

    var i1 = 5
    fmt.Printf("An integer: %d, it location in memory: %p\n", i1,&i1)
    // An integer: 5, it location in memory: 0x42133120

    var intP * int
    intP = &i1
    fmt.Printf("The value at memory location %p is %d\n", intP, *intP)
    // The value at memory location 0x42133120 is 5
}
