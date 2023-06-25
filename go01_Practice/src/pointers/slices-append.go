package main

import "fmt"

func main() {
	var s []int
	printSlice12(s)

	// 添加一个空切片
	s = append(s, 0)
	printSlice12(s)

	// 这个切片会按需增长
	s = append(s, 1)
	printSlice12(s)

	// 可以一次性添加多个元素
	s = append(s, 2, 3, 4)
	printSlice12(s)
}

func printSlice12(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
