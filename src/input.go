// Input a float number and double it
package main

import "fmt"

func main() {
  fmt.Print("Input a number: ")
  var input float64
  fmt.Scanf("%f", &input)
  output := input * 2
  fmt.Println(output)
}


