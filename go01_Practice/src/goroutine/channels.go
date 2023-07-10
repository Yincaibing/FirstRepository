package main

import "fmt"
/*
信道是带有类型的管道，你可以通过它用信道操作符 <- 来发送或者接收值。
ch <- v    // 将 v 发送至信道 ch。
v := <-ch  // 从 ch 接收值并赋予 v。
（“箭头”就是数据流的方向。）

和映射与切片一样，信道在使用前必须创建：
ch := make(chan int)
默认情况下，发送和接收操作在另一端准备好之前都会阻塞。这使得 Go 程可以在没有显式的锁或竞态变量的情况下进行同步。

在Go语言中，Channel是一种特殊类型的数据结构。这是Go语言对生产者和消费者模型实现的一种方式，它允许一个goroutine向另一个goroutine发送数据。Channel在goroutine之间提供了一种通信机制，所以能很简单地通过channel在goroutines里进行数据交换。
设计理念：Go语言的并发编程模型主要采用CSP（communicating sequential processing）理论，而channel正是CSP理论的核心部分。CSP理论指出在并发系统中，最重要的并不是并发的实体（线程或进程），而是实体之间的通信。Go的设计者在设计语言时引入了基于CSP的goroutine和channel概念。
channel的作用：
1. 线程安全：多个goroutine可以同时往同一个channel写数据或者同一个channel读数据。
2. 数据交换：在两个goroutine间安全交换数据，避免共享数据带来的复杂问题。
3. 同步与阻塞：不同于其他多线程并发语言使用锁的实现，Go语言的channel提供了一种新颖的同步方式。在读取或者写入channel时，如果对应的channel无法完成读或者写的操作，这个操作就会阻塞，直到可以完成该操作。
4. 管道：可以把多个channel链接起来，一个channel的输出作为下一个channel的输入，形成管道（pipeline）模式。
通过以上所述，你应该对Go语言中Channel的设计理念和作用有了更深入的理解。

以下示例对切片中的数进行求和，将任务分配给两个 Go 程。一旦两个 Go 程完成了它们的计算，它就能算出最终的结果。
*/
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将和送入 c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	//制作一个channel的数据结构用来存放结果，具有同步的作用
	c := make(chan int)

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从 c 中接收

	fmt.Println(x, y, x+y)
}
