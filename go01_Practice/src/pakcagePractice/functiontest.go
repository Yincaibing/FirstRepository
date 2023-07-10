package main

//Go 的声明从左到右阅读
//函数
//函数可以没有参数或接受多个参数，也可以返回多个参数
//在本例中，add 接受两个 int 类型的参数。
//注意类型在变量名 之后。
//（参考 这篇关于 Go 语法声明的文章了解这种类型声明形式出现的原因。https://blog.go-zh.org/gos-declaration-syntax）

/*
func funcName(input1 type1, input2 type2) (output1 type1, output2 type2) {
	//这里是处理逻辑代码
	//返回多个值
	return value1, value2
}

关键字func用来声明一个函数funcName
函数可以有一个或者多个参数，每个参数后面带有类型，通过,分隔
函数可以返回多个值
上面返回值声明了两个变量output1和output2，如果你不想声明也可以，直接就两个类型
如果只有一个返回值且不声明返回值变量，那么你可以省略 包括返回值 的括号
如果没有返回值，那么就直接省略最后的返回信息
如果有返回值， 那么必须在函数的外层添加return语句
*/

import "fmt"

func add(x int, y int) int {
	return x + y
}

// 当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略。
// 在本例中，x int, y int被缩写为x, y int
func reduce(x, y int) int {
	return x + y
}

// 函数可以返回任意数量的返回值。这里返回了两个字符串
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(reduce(13, 12))
	a, b := swap("AAA", "BBB")
	fmt.Println(a, b)

}
