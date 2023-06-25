package main

//Go 的声明从左到右阅读
//函数
//函数可以没有参数或接受多个参数，也可以返回多个参数
//在本例中，add 接受两个 int 类型的参数。
//注意类型在变量名 之后。
//（参考 这篇关于 Go 语法声明的文章了解这种类型声明形式出现的原因。https://blog.go-zh.org/gos-declaration-syntax）

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
