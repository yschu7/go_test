package main

import ("fmt"; "math")

type Circle struct { x, y, r float64 }
type Rectangle struct { x1, y1, x2, y2 float64 }

type Shape interface { area() float64 }
func totalArea(shapes ...Shape) float64 {
    var area float64
    for _, s := range shapes {
        area += s.area()
    }
    return area
}

type MultiShape struct { shapes []Shape }
func (m *MultiShape) area() float64 {
    var area float64
    for _, s := range m.shapes {
        area += s.area()
    }
    return area
}

func distance(x1,y1,x2,y2 float64) float64 {
  a := x2 - x1
  b := y2 - y1
  return math.Sqrt(a*a + b*b)
}
// method area() of Rectangle
func (rec *Rectangle) area() float64 {
    l := distance(rec.x1, rec.y1, rec.x1, rec.y2)
    w := distance(rec.x1, rec.y1, rec.x2, rec.y1)
    return l * w
}
// method area() of Circle
func (cir *Circle) area() float64 {
    return math.Pi * cir.r * cir.r
}
func main() {
    rec := Rectangle{0, 0, 10, 10}
    cir := Circle{0, 0, 5}

    fmt.Println(rec.area()) // call method with "." operator
    fmt.Println(cir.area())
    fmt.Println(totalArea(&rec, &cir))
    mul := new(MultiShape)
    mul.shapes = make([]Shape,2) // allocate space for slice
    mul.shapes[0] = &rec
    mul.shapes[1] = &cir
    fmt.Println(mul.area())
}

