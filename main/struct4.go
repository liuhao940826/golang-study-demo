package main

import "fmt"

type Student3 struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

//结构体指针变量初始化
func main() {
	//指针有合法指向后才操作成员

	//先定义一个普通结构体对象

	var s Student3

	//在定义一个指针变量在保存s的地址
	var p1 *Student3
	p1 = &s
	//p1.id和 *p1.id 完全等价 只能使用.运算符
	p1.id = 18
	(*p1).name = "mike"
	s.sex = 'm'
	s.addr = "hahah"
	s.age = 18
	fmt.Println("s = ", s)

	//通过new来申请结构体 p2 是指针类型
	p2 := new(Student3)

	fmt.Printf("p2的类型 %T \n", p2)
	p2.id = 18
	p2.name = "mike"
	p2.sex = 'm'
	p2.addr = "hahah"
	p2.age = 18
	fmt.Println("p2 = ", p2)

}
