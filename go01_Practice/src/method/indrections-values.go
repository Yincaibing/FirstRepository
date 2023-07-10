package main

import (
	"fmt"
	"math"
)

type Vertex111 struct {
	X, Y float64
}

func (v Vertex111) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex111) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex111{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex111{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}
