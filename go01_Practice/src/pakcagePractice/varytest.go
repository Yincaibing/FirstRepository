package main

//1、变量的声明
//var 语句用于声明一个变量列表，跟函数的参数列表一样，类型在最后。
//就像在这个例子中看到的一样，var 语句可以出现在【包】或【函数】级别。

//2、变量的初始化
//变量声明可以包含初始值，每个变量对应一个。
//如果初始化值已存在，则可以省略类型；变量会从初始值中获得类型。
import "fmt"

// 1、声明一个自定义变量，类型为bool类型
var c, python, java bool

// 2、变量声明可以包含初始值，每个变量对应一个
var i, j int = 1, 2

func main() {
	// 1、声明一个自定义变量，类型为int整型
	var i int
	fmt.Println(i, c, python, java)

	//2、声明变量并赋不同类型的初始值，
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
