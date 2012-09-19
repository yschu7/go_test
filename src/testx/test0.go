package main
import (
   "fmt"
   "./trans"
)
var globalVar int
var twoPi = 2 * trans.Pi

func main() {
   //globalVar = 33
   fmt.Printf("2*Pi = %g\n", twoPi)

   hello()        // test1.go
   printVar()     // test2.go
   printConst()   // test2.go
   fmt.Println(變)    // global variable declared in test2.go
}

