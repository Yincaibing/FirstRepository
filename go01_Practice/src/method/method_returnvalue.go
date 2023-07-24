package main


/*
分别介绍方法的返回值用法，以及函数式编程

*/
func main() {
	Invoke()
}



func Invoke() {
	str := Func0("Fun0")
	println(str)

	func1, err := Func1(1, 2, 3, "Fun1")
	if err != nil {
		return 
	}
	println(func1)

	
	func2, err := Func2(1, 2)
	if err != nil {
		return 
	}
	println(func2)

	
	func3, err := Func3(1, 2)
	if err != nil {
		return 
	}
	println(func3)

	Func4()

	Func5()

	sayHello:=Func6()
	sayHello("caibing,this is Func6, the Func as a return value")
	
}

// Func6 函数式编程，因为方法本身是可以作为返回值的，方法本身就是一个代码块，他本身就是播报的一种应用。
func Func6() func(name string) string {
	return func(name string) string {
		return "hello,"+name
	}
}

// Func5 可以在方法内部声明一个方法，他的作用域就在方法内
func Func5(){
	fn := func(name string) string {
		return "hello, "+ name
	}
	fn("caibing")
}

// Func4 这里面的 Func3本质上是一个变量。只不过这个变量等于 Func3； 因此，在这里有一个新的用法：函数式编程
func Func4() {
	myFunc3 := Func3
	myFunc3(1,2)
}


// Func3 带名字的返回值，虽然都取了名字，返回的时候，可以只接受一个，另一个做默认值处理，可能另一个也不会用到
func Func3(i int, i2 int)(str string, err error) {
	res :="hello,Func3"
	return res,nil
}

// Func2 带名字的返回值，可以直接 return;并且如果要带名字，两个都必须要带名字，不能一个带，一个不带
func Func2(a,b int)(str1 string,err error) {
	str1 = "hello,Func2"
	return 
}

// Func1 带2个返回值
func Func1(a,b,c int, s string)(string,error) {
	return s,nil
}

// Func0 单一返回值
func Func0(name string)string   {
	return "hello,"+name
}

