// How to use for loop
package main

import "fmt"

func main() {
  i := 1
  for i <= 10 {
    fmt.Println(i)
    i += 1
  }
  for i := 1; i <= 5; i++ {
    fmt.Println(i)
  }
}

