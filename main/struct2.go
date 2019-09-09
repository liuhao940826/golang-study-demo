package main

import "fmt"

type Student1 struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

//结构体指针变量初始化
func main() {

	//顺序初始化 指针变量
	var p1 *Student1 = &Student1{1, "s1", 'm', 18, "sh"}
	fmt.Println("*p1", *p1)

	//指定成员初始化,没有初始化的成员,自动赋值为0
	p2 := &Student1{name: "mike", addr: "上海"}
	fmt.Printf("p2  type is %T \n", p2)
	fmt.Println("*p2", *p2)

}
