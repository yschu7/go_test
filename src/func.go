package main
import (
  "fmt"
)
func add(args ...int) int {
    total := 0
    for _, v := range args {
        total += v
    }
    return total
}
func main() {
    fmt.Println(add(1,2,3))
    xs := []int{3,4,5}
    fmt.Println(add(xs...))
    nextEven := makeEvenGenerator()
    fmt.Println(nextEven()) // 0
    fmt.Println(nextEven()) // 2
    fmt.Println(nextEven()) // 4
}
func makeEvenGenerator() func() uint {
    i := uint(0)
    return func() (ret uint) {
        ret = i
        i += 2
        return
    }
}
