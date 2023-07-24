package main

/*
1、匿名方法：
在方法内部可以声明一个匿名方法，但是需要立即发起调用。为什么？
因为匿名，如果不立刻调用的话后面没办法调用了。匿名方法在 defer中用的多

2、闭包：
方法+他绑定的运行上下文


3、不定参数
Go 支持不定参数。不定参数是指最后一个参数，可以传入任意多个值。
注意必须是最后一个参数才可以声明为不定参 数。不定参数在方法内部可以被当成切片来使用。
Option 模式大量应用了不定参数。



*/
func main() {
	Func7()
	YourName("YCB0")
	YourName("YCB0","YCB1","YCB2","YCB3")
}

// Func7
func Func7() {
	hello := func()string{
		return "hello,caibing, We are testing 匿名函数"
	}
	println(hello)
}

// Closure 闭包如果使用不当可能会引起内存泄漏的问题。即一个对象被闭包引用的话，他是不会被垃圾回收的。
func Closure(name string) func() string{
	//返回的这个函数就是一个闭包
	//他引用了 Closure这个方法的入参
	return func() string {
		return "hello, " + name
	}
}


// YourName 不定参数方法，这个方法入参有多个参数，最后一个参数可以声明为不定参数
func YourName(name string, alias ...string){
	if len(alias)>0{
		println("this is second param: "+ alias[1])
	}
}
