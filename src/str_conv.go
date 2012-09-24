package main

import (
	"fmt"
	"strconv"
//  "os"
)

func main() {
	var orig string = "666"
	var an int
	var newS string
  // var err error
	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
  /*
	an, err = strconv.Atoi(orig)
  if err != nil {
    fmt.Printf("An error occured in strconv.Atoi with %v\n",orig)
    os.Exit(1)
  }
  */
  an = atoi(orig)
	fmt.Printf("The Integer is: %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)
	str1 := strconv.FormatInt(int64(an), 2)
	str2 := strconv.FormatInt(int64(an), 16)
	fmt.Println("binary: ", str1, " hex: ", str2)
	var sb []byte
	sb = strconv.AppendInt(sb, 5, 2)
	fmt.Println(sb)
	fmt.Println(string(sb))
	fmt.Println(strconv.Quote(string(sb)))
}

func atoi(s string)(n int) {
  n, _ = strconv.Atoi(s)
  return
}
