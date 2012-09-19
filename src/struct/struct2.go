package main

import ("fmt"; "math")

type Circle struct { x, y, r float64 }
type Rectangle struct { x1, y1, x2, y2 float64 }

func distance(x1,y1,x2,y2 float64) float64 {
  a := x2 - x1
  b := y2 - y1
  return math.Sqrt(a*a + b*b)
}
func rectangleArea(rec *Rectangle) float64 {
    l := distance(rec.x1, rec.y1, rec.x1, rec.y2)
    w := distance(rec.x1, rec.y1, rec.x2, rec.y1)
    return l * w
}
func circleArea(cir *Circle) float64 {
    return math.Pi * cir.r * cir.r
}
func main() {
    rec := Rectangle{0, 0, 10, 10}
    cir := Circle{0, 0, 5}

    fmt.Println(rectangleArea(&rec))
    fmt.Println(circleArea(&cir))
}

