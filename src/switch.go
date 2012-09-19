package main
import "fmt"
func main() {
  i := 2
  str := `this is a
multiline string
test`
  str1 := "this is a\nsingle line string\ntest"
  var f float32 = 5.0
  fmt.Println(f)
  switch i {
    case 0: fmt.Println("Zero")
    case 1: fmt.Println("One")
    case 2: fmt.Println("Two")
      fmt.Println(str)
      fmt.Println(str1)
    case 3: 
      fmt.Println("Three")
      fmt.Println("Three....")
      fmt.Println("Three.x...")
    case 4: fmt.Println("Four")
    case 5: fmt.Println("Five")
    default: fmt.Println("Unknown Number")
  }
}
