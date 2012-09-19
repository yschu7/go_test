package main

import ("fmt";"flag";"math/rand";"time")

func main() {
    // Define flags
    maxp := flag.Int("max", 6, "the max value")
    // Parse
    flag.Parse()
    // Generate a number between 0 and max
    rand.Seed(time.Now().Unix())
    fmt.Println(rand.Intn(*maxp))
}

