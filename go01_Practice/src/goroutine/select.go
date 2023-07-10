package main

import "fmt"
/*
select 语句使一个 Go 程可以等待多个通信操作。
select 会阻塞到某个分支可以继续执行为止，这时就会执行该分支。当多个分支都准备好时会随机选择一个执行。
select 语句在 Go 语言编程中广泛用于实现并发程序，通过它可以轻松地处理从不同 channel 接收的数据，有助于编写清晰、易于理解的并发程序。
*/

//这个函数需要传入两个整型数据的 chan
func fibonacci12(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		//并发逻辑 1：x进入chan c后，进行x，y的操作
		case c <- x:
			x, y = y, x+y
		//并发逻辑 2：从 quit chan出来的数据，由于没有任何操作，表示可以结束了
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci12(c, quit)
}
