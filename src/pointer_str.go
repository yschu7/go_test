package main
import (
  "fmt"
)
func main() {
	s := "good bye"
	var p *string = &s
	*p = "ciao"
	fmt.Printf("pointer p: %p\n", p)
	fmt.Printf("string *p: %s\n", *p)
	fmt.Printf("string  s: %s\n", s)

	const i = 5
	prt := &i  // error: cannot take the address of i
}
