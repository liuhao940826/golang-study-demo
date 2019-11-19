package main

import "fmt"

type Student struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

//数组长度在定义后无法再次修改 数组是值类型 每次传递产生一个副本  用slice来弥补不足
func main() {

	s := Student{}
	fmt.Printf("默认值%+v \n", s)

	//顺序初始化
	var s1 Student = Student{1, "s1", 'm', 18, "sh"}
	fmt.Println("s1", s1)

	//指定成员初始化,没有初始化的成员,自动赋值为0
	s2 := Student{name: "mike", addr: "上海"}
	fmt.Println("s2", s2)

}
