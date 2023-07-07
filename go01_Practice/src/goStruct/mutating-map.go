package main

func main() {
	m:= make(map[string]int)

	m["key1"]=12
	println("the value is: ", m["key1"])

	m["key1"]=13
	println("the value is: ", m["key2"])

	// 获取元素：value = m[key]
	m["Answer"] = 48
	println("The value:", m["Answer"])

	//删除元素：delete(m, key)
	delete(m, "Answer")
	println("The value:", m["Answer"])

	//通过双赋值检测某个键是否存在：在 go中经常用到，非常高频。
	//注 ：若 elem 或 ok 还未声明，你可以使用短变量声明：
	v, ok := m["Answer"]
	println("The value:", v, "Present?", ok)

}
