package main

import "fmt"


//nil 接口值
//nil 接口值既不保存值也不保存具体类型。
//为 nil 接口调用方法会产生运行时错误，因为接口的元组内并未包含能够指明该调用哪个 具体 方法的类型。

type I12 interface {
	M()
}

func main() {
	var i I12
	describe12(i)
	//这里接口没有实现，所以接口值为 nil，
	i.M()
}

func describe12(i I12) {
	fmt.Printf("(%v, %T)\n", i, i)
}
