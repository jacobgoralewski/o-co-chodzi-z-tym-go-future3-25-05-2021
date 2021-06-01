package main

import "fmt"

type Vertex struct {
	x int
	y int
}

func (v Vertex) SetX(x int) {
	v.x = x
}

func (v *Vertex) PointerSetX(x int) {
	v.x = x
}

func main() {
	v := Vertex{1, 2}
	v.SetX(3)
	fmt.Println(v)

	p := &v
	p.PointerSetX(3)
	fmt.Println(p)
}