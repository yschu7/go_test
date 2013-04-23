package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
    // Generate a number between 0 and max
		max := 10
    rand.Seed(time.Now().Unix())
    for i := 0; i < 6; i++ {
	    fmt.Print(rand.Intn(max), " / ")
	  }
	  fmt.Println()
    // float
    times := int64(time.Now().Nanosecond())
    rand.Seed(times)
    for i := 0; i < 6; i++ {
	    fmt.Printf("%2.2f / ", 100*rand.Float32())
	  }
}

//
