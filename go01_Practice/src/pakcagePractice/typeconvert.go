package main

import (
	"fmt"
	"math"
)

//1、【类型转换】
// 表达式 T(v) 将值 v 转换为类型 T。

func main() {
	//下面是一些数值转换的例子
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	//或者，更加简单的形式：
	ii := 42
	ff := float64(ii)
	uu := uint(ff)
	fmt.Println(ii, ff, uu)

}
