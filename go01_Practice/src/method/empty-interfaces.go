package main

import "fmt"

//指定了零个方法的接口值被称为 *空接口：*
//interface{}
//空接口可保存任何类型的值，因为每个类型都至少实现了零个方法。
//空接口用法：
//	空接口被用来处理【未知类型】的值。例如，fmt.Print 可接受类型为 interface{} 的任意数量的参数。
/*
空 interface的最大用处是：一个函数把interface{}作为参数，那么他可以接受任意类型的值作为参数，如果一个函数返回interface{},那么也就可以返回任意类型的值。是不是很有用啊！

// 定义a为空接口
var a interface{}
var i int = 5
s := "Hello world"
// a可以存储任意类型的数值
a = i
a = s

*/

func main() {
	//这个 i接口没有实现任何方法
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
