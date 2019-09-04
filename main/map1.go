package main

import "fmt"

func main() {

	var m1 map[int]string

	fmt.Println("m1=", m1)
	//对于map只有len() 没有cap
	fmt.Println("len=", len(m1))

	//通过make来创建  这个是make创建 片 s:= make([]int, 0, 1)
	m2 := make(map[int]string)
	fmt.Println("m2=", m2)
	fmt.Println("len=", len(m2))

	//通过make来创建  可以指定长度,只是指定了map的容量 但是里面一个数据都没有 这个容量可以扩充的
	m3 := make(map[int]string, 2)
	m3[1] = "mike"
	m3[2] = "go"
	m3[3] = "c++"
	fmt.Println("m3=", m3)
	fmt.Println("len=", len(m3))

	m4 := map[int]string{1: "mike", 2: "go", 3: "c++"}
	fmt.Println("m4=", m4)

}
