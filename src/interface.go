package main

type Concrete int

func (c Concrete) Foo() {
    println("concrete foo")
}

func (c Concrete) Bar() {
    println("concrete bar")
}

type Interface1 interface {
    Foo()
    Bar()
}

type Interface2 interface {
    Foo()
}

type MyWrapper struct {
    Interface2
}

func (my MyWrapper) Bar() {
    println("my wrapper bar")
}

func main() {
    var c Concrete
    var i1 Interface1 = c
    var x Interface2
    y := MyWrapper {x}
    i1.Foo() // "concrete foo"
    i1.Bar() // "concrete bar"
    y.Bar()  // "my wrapper bar"
}