package main

import "fmt"

func main() {
	//1、跟C或者java相比，for循环条件表达式没有小括号，其他一样
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//2、初始化语句和后置语句是可选的。
	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	//3、for 是 Go 中的 “while”   此时你可以去掉分号，因为 C 的 while 在 Go 中叫做 for。
	sum3 := 1
	for sum3 < 1000 {
		sum3 += sum3
	}
	fmt.Println(sum3)

	//4、无限循环
	//如果省略循环条件，该循环就不会结束，因此无限循环可以写得很紧凑。
	for {
	}
}
