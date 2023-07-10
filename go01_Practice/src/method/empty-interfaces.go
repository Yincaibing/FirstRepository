package main

import "fmt"

//指定了零个方法的接口值被称为 *空接口：*
//interface{}
//空接口可保存任何类型的值，因为每个类型都至少实现了零个方法。
//空接口用法：
//	空接口被用来处理【未知类型】的值。例如，fmt.Print 可接受类型为 interface{} 的任意数量的参数。

func main() {
	//这个 i接口没有指定任何方法
	var i interface{}
	describe123(i)

	//用来处理未知类型的值
	i = 42
	describe123(i)

	i = "hello"
	describe123(i)
}

//接口作为参数
func describe123(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
