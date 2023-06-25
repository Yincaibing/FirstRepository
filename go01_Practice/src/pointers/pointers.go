package main

import "fmt"

//& 操作符会生成一个指向其操作数的指针。  p := &i
//* 操作符表示指针指向的底层值。  	fmt.Println(*p) // 通过指针读取 i 的值，为 42
//这也就是通常所说的“间接引用”或“重定向”。
//与 C 不同，Go 没有指针运算。

func main() {
	i, j := 42, 2701

	p := &i         // 指向 i
	fmt.Println(*p) // 通过指针读取 i 的值，为 42
	*p = 21         // 通过指针设置 i 的值
	fmt.Println(i)  // 查看 i 的值，为 21

	p = &j         // 指向 j
	*p = *p / 37   // 通过指针对 j 进行除法运算2701/37=
	fmt.Println(j) // 查看 j 的值，为 73
}
