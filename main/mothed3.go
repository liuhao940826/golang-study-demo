package main

import (
	"fmt"
)

type Person6 struct {
	name string
	sex  byte
	age  int
}

func (tmp Person6) PrintInfo2() {
	fmt.Println("tmp=", tmp)
}

//通过一个函数,给成员赋值  这个Person6 本身不能够是指针类型
func (tmp *Person6) SetInfo2(name string, sex byte, age int) {

	tmp.name = name
	tmp.age = age
	tmp.sex = sex
}

func main() {

}
