package main

import "fmt"

type Vertex12 struct {
	Lat, Long float64
}

var m12 = map[string]Vertex12{
	"Rock Labs": Vertex12{40.68433, -74.39967},   //如果Vertex12只是一个类型，可以省略不写
	"Bell Labs": {40.68433, -74.39967},   //如果Vertex12只是一个类型，可以省略不写
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m12)
}
