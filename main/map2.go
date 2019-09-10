package main

import "fmt"

func main() {

	//map的常用 操作  赋值
	m4 := map[int]string{1: "mike", 2: "yoyo", 3: "c++"}
	fmt.Println("m4=", m4)
	//如果已经存在的key  赋值会覆盖
	m4[1] = "lala"
	fmt.Println("m4=", m4)
	//key不存在会直接追加
	m4[3] = "yien"
	fmt.Println("m4=", m4)

}
