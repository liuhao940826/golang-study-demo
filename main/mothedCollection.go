package main

import (
	"fmt"
)

type Person8 struct {
	name string
	sex  byte
	age  int
}

//接受者为普通变量 值传递
func (tmp Person8) SetValueInfo() {
	//tmp.name = "mike"
	//tmp.age = 18
	//tmp.sex = 'f'
	fmt.Println("setValue")
}

//接受者为指针变量 引用传递
func (tmp *Person8) SetPointerInfo() {
	//
	//tmp.name = name
	//tmp.age = age
	//tmp.sex = sex
	fmt.Println("SetPointerInfo")
}

func main() {
	//加入这个结构体变量是一个指针变量 他能够调用哪些方法,这些方法就是一个集合 ,简称方法集
	p := &Person8{"mike", 'm', 18}
	p.SetPointerInfo()
	//内部做了转换 把指针变量p 转黄成了*p在调用 同理反过来 如果p不是指针变量 就是Person8也一样
	p.SetValueInfo()

}
