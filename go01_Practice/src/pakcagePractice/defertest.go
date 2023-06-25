package main

//defer

import (
	"fmt"
	"io"
	"os"
)

//3、defer常常用在程序最后的操作等，比如忘记关掉文件符

// 有bug,如果异常，可能会忘了关掉文件句柄
func CopyFileHasbug(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}

	written, err = io.Copy(dst, src)
	dst.Close()
	src.Close()
	return
}

// 用了defer之后，如果有异常，最后一步关掉文件句柄
func CopyFileNoBug(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}
func main() {
	//1、defer 语句会将函数推迟到外层函数返回之后执行。
	//推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。
	defer fmt.Println("world")
	fmt.Println("hello")

	//2、defer 栈
	//推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。
	//更多关于 defer 语句的信息，请阅读此博文：https://blog.go-zh.org/defer-panic-and-recover
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")

}
