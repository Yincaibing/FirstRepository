package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat1234(-math.Sqrt2)
	v := Vertex12345{3, 4}

	a = f  // a MyFloat 实现了 Abser
	a = &v // a *Vertex 实现了 Abser

	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	// 所以没有实现 Abser。
	a = v

	fmt.Println(a.Abs())
}

type MyFloat1234 float64

func (f MyFloat1234) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex12345 struct {
	X, Y float64
}

func (v *Vertex12345) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
