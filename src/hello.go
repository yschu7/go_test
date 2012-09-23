#!/usr/bin/env gorun
package main

import (
  fm "fmt"
//  "./test.go"
)
func main() {
  n := -1
  fm.Printf("%d > 0 : %t \n", n, isTrue(n))
}
func isTrue(x int) bool{
  if x > 0 {
    return true
  }
  return false
}
