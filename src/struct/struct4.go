package main
import (
  "fmt"
)
type Person struct {
    Name string
}
func (p *Person) Talk() {
    fmt.Println("Hi, my name is", p.Name)
}
type Android struct {
    Person
    Model string
}
func main() {
  a := new(Android)
  a.Person.Name = "Robot"
  a.Person.Talk()
  b := Android{ Person{"Number 9"}, "P-41"}
  b.Talk()
}
