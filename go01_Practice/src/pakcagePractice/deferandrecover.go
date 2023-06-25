package main

import "fmt"

//This is convenient for modifying the error return value of a function; we will see an example of this shortly.
//Panic is a built-in function that stops the ordinary flow of control and begins panicking. When the function F calls panic, execution of F stops, any deferred functions in F are executed normally, and then F returns to its caller.To the caller, F then behaves like a call to panic. The process continues up the stack until all functions in the current goroutine have returned, at which point the program crashes. Panics can be initiated by invoking panic directly.They can also be caused by runtime errors, such as out-of-bounds array accesses.
//Recover is a built-in function that regains control of a panicking goroutine. Recover is only useful inside deferred functions. During normal execution, a call to recover will return nil and have no other effect. If the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution.
//
//Here's an example program that demonstrates the mechanics of panic and defer:

//这样方便修改函数的错误返回值；我们将会看到这样的例子。
//Panic 是一个内置函数，它停止普通的控制流并开始恐慌。当函数 F 调用 panic 时，F 的执行停止，F 中的任何延迟函数都正常执行，然后 F 返回到它的调用者。对于调用者，F 然后表现得像调用 panic。该过程继续向上堆栈，直到当前 goroutine 中的所有函数都已返回，此时程序崩溃。恐慌可以通过直接调用恐慌来启动。它们也可能是由运行时错误引起的，例如越界数组访问。
//Recover 是一个内置函数，可以重新获得对 panic 协程的控制。恢复仅在延迟函数内部有用。在正常执行期间，对 recover 的调用将返回 nil 并且没有其他影响。如果当前 goroutine 处于恐慌状态，则调用 recover 将捕获给予恐慌的值并恢复正常执行。

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
