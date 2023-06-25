package main

import "fmt"


/**
map
map的零值为 nil 。nil map既没有键，也不能添加键。
make 函数会返回给定类型的map，并将其初始化备用。
*/


//先定义一个键值对的结构体

type Vertex struct {
	Lat, Long float64
}

//再创建一个该类型的数组
var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,   //给该键值对赋值
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(len(m))    //该长度只有 1
}