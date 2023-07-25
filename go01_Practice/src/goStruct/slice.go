package main

func main() {
	s1 := make([]int,3,4)    //直接初始化了三个元素，容量为 4 的切片
	for len(s1)!=0{
		println(cap(s1))
		break
	}


	//方法 2：sliceName := []dataType{element1, element2, ..., elementN}
	s2 := []int{1,2,3,4}
	println(len(s2))
	println(cap(s2))
	
	s2 = append(s2, 5)   //追加了一个元素
	println(len(s2))
	println(cap(s2))


}
