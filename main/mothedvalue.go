package main

import (
	"fmt"
)

type Person9 struct {
	name string
	sex  byte
	age  int
}

//接受者为普通变量 值传递
func (tmp Person9) SetValueInfo() {
	fmt.Println("setValue")
}

//接受者为指针变量 引用传递
func (tmp *Person9) SetPointerInfo() {
	fmt.Println("SetPointerInfo")
}

func main() {
	//加入这个结构体变量是一个指针变量 他能够调用哪些方法,这些方法就是一个集合 ,简称方法集
	p := Person9{"mike", 'm', 18}
	fmt.Printf("main %p ,%+v", &p, p)

	p.SetValueInfo() //传统调用方式

	//保存函数的入口
	pFunc := p.SetPointerInfo //这个pFunc就是方法值 调用函数式不需要在用调用着了
	pFunc()

}
