package main

import "fmt"

type Student2 struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

//结构体指针变量初始化
func main() {

	//定义一个结构体普通变量
	var s Student2
	//操作成员需要使用. 运算符
	s.id = 1
	s.name = "aaa"
	s.sex = 'm'
	s.addr = "hahah"
	s.age = 18

	fmt.Println("s = ", s)

}
