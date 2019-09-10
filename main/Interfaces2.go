package main

import (
	"fmt"
)

type Humaner1 interface {
	SayHi() //这个叫方法 只有声明 没有实现 由别的(自定义类型)实现
}

type Student11 struct {
	name string
	id   int
}

type Teacher12 struct {
	addr  string
	group string
}

type MyStr1 string

func (tmp *Student11) SayHi() {
	fmt.Printf("Student[%s ,%d] sayHi \n", tmp.name, tmp.id)
}

func (tmp *Teacher12) SayHi() {
	fmt.Printf("Teacher11[%s ,%s] sayHi \n", tmp.addr, tmp.group)
}

func (tmp *MyStr1) SayHi() {
	fmt.Printf("MyStr[%s] sayHi \n", *tmp)
}

//定义一个普通函数

func WhoSayHi(i Humaner1) {
	i.SayHi()
}

func main() {
	fmt.Println("interfaces")

	//只要实现了此接口中所有方法的类型,那么这个类型的变量(接收类型)就可以给i赋值

	//实现接口
	s := &Student11{"liuhao", 666}

	//调用同一个函数不同表现
	WhoSayHi(s)

	t := &Teacher12{"bj", "go"}
	//调用同一个函数不同表现
	WhoSayHi(t)

	var str MyStr1 = "hello world"
	//调用同一个函数不同表现
	WhoSayHi(&str)
	fmt.Println("=======================")
	//创建一个分片 类型 []humaner   格式make(切片类型, 长度,cap)
	x := make([]Humaner1, 3)

	x[0] = s
	x[1] = t
	x[2] = &str
	//第一个返回下标 第二个返回下标对应的值
	for _, j := range x {
		j.SayHi()
	}

}
