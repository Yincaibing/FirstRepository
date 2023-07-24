package main

import "fmt"

//【短变量声明】
//     k := 3，这种声明只限于方法内部，并且是go支持的类型推断，数字会被理解为 int或者 float64.(所以要其他类型的数字，就得用 var来声明)
//在函数中，简洁赋值语句 := 可在类型明确的地方代替 var 声明。
//函数外的每个语句都必须以关键字开始（var, func 等等），因此 := 结构不能在函数外使用。

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
}
