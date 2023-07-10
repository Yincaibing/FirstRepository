package main

import (
	"fmt"
	"math"
)

type Vertex123 struct {
	X, Y float64
}

func (v Vertex123) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//将 Abs转化为函数
func AbsFunction(v Vertex123) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}


//通过指针操作原来的结构体
func (v *Vertex123) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//将 Scale转化为函数
func ScaleFunction(v *Vertex123, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex123{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())

	//函数调用方式跟方法也不一样
	ScaleFunction( &v,5)


}
