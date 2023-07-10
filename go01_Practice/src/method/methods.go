package main

import (
	"fmt"
	"math"
)

/**
Go语言是一种集功能、并发和面向对象编程的编程语言。但是，与传统的面向对象语言如Java或C++不同，
Go语言确实没有"类"的概念。这主要是因为Go语言设计者认为传统的类层次结构（尤其是使用继承的那一部分）是源码的脆弱部分，它们更喜欢更简单、更直接的组合和接口。
尽管没有显性的"类"概念，但Go语言还是支持面向对象的编程方式，它使用结构体（struct）和方法（method）来实现。
Go语言鼓励使用【组合代替继承】，通过【为结构体添加方法来实现面向对象】的特性。通过结构体来组合数据和行为，可以灵活地构建和管理对象的行为

注：方法与函数的区别：
参考这里，关于面向对象这里讲的很通透
https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.5.md

method的语法如下：
func (r ReceiverType) funcName(parameters) (results)
*/

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {    //方法接收者在它自己的参数列表内，位于 func 关键字和方法名之间。这是不同于函数的地方，他们有自己的接受者，Abs方法依赖于 Vertex结构体对象来接受他
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type car struct {
	Brand  string
	Color  string
	Model  string
	Speed  int
}


type vehicle interface {    //定义一个接口
	Drive(speed int)
}


func (c car) Drive(speed int) {
	c.Speed = speed
	fmt.Printf("%s %s is now driving at %d km/h.\n", c.Color, c.Brand, c.Speed)
}


func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	
	lexus := car{
		"lexus","red","ES300h",120,
	}
	lexus.Drive(120)
}
