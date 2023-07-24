package main

/*
iota 可以方便地控制常量初始化。
可以使用 iota，也可以使用 iota 的数学运 算，包括位移操作。
只有主动赋值或者另起一个 iota 才会从头 开始计算值。
*/

func main() {
	const(
		Status0 =iota   //注意这里，只需要声明 Status0=iota，即可自动加1，这样的话，后面变量不用显示声明，也能知道自动加一，除非显示声明主动赋值，或者令其一个 iota才会从头开始计算值
		Status1
		Status2
		Status3


		Status6=6  //中断 iota赋值
	)

	const (
		One =iota<<1
		Two
		Four
		) 

}