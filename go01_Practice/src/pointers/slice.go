package main

import "fmt"

/**
切片
每个数组的大小都是固定的。而切片则为数组元素提供动态大小的、灵活的视角。在实践中，切片比数组更常用。
类型 []T 表示一个元素类型为 T 的切片。
切片通过两个下标来界定，即一个上界和一个下界，上界是从 0 开始计数的，二者以冒号分隔：
a[low : high]
它会选择一个半开区间，包括第一个元素，但排除最后一个元素。
以下表达式创建了一个切片，它包含 a 中下标从 1 到 3 的元素：
a[1:4]
*/


func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] //[3，,5，,7】]
	fmt.Println(s)

	//切片就像数组的引用
	//切片并不存储任何数据，它只是描述了底层数组中的一段。
	//更改切片的元素会修改其底层数组中对应的元素。
	//与它共享底层数组的切片都会观测到这些修改。

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2] //["John","Paul"]
	b := names[1:3] //["Paul","George"]
	fmt.Println(a, b)

	b[0] = "XXX" //["XXX","George"]
	fmt.Println(a, b)
	fmt.Println(names)
}
