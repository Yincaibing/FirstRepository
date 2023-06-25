package main

import "fmt"

//结构体字段使用点号来访问。

type Vertex1 struct {
	X int
	Y int
}

func main() {
	v := Vertex1{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
