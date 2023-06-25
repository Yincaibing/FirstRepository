package main

import "fmt"

//1、【常量】
// 常量的声明与变量类似，只不过是使用 const 关键字。
// 常量可以是字符、字符串、布尔值或数值。
// 常量不能用 := 语法声明。

const Pi = 3.14

//2、【数值常量】
//数值常量是高精度的 值。
//一个未指定类型的常量由上下文来决定其类型。
//再尝试一下输出 needInt(Big) 吧。
//（int 类型最大可以存储一个 64 位的整数，有时会更小。）
//（int 可以存放最大64位的整数，根据平台不同有时会更少。）

const (
	// Big 将 1 左移 100 位来创建一个非常大的数字
	// 即这个数的二进制是 1 后面跟着 100 个 0
	Big = 1 << 100
	// Small 再往右移 99 位，即 Small = 1 << 1，或者说 Small = 2
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
