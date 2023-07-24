package main

func main() {
	var a byte = 'a'
	println(a)   //输出的是 a的 ascII表达 97
	
	var str string = "this is string"
	var bs []byte = []byte(str)    //类型转换
	println(bs)
}
