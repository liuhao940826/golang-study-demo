package main

import (
	"fmt"
)

type Humaner interface {
	SayHi() //这个叫方法 只有声明 没有实现 由别的(自定义类型)实现
}

type Student10 struct {
	name string
	id   int
}

type Teacher11 struct {
	addr  string
	group string
}

type MyStr string

func (tmp *Student10) SayHi() {
	fmt.Printf("Student[%s ,%d] sayHi \n", tmp.name, tmp.id)
}

func (tmp *Teacher11) SayHi() {
	fmt.Printf("Teacher11[%s ,%s] sayHi \n", tmp.addr, tmp.group)
}

func (tmp *MyStr) SayHi() {
	fmt.Printf("MyStr[%s] sayHi \n", *tmp)
}

func main() {
	fmt.Println("interfaces")
	//定义一个接口类型的变量

	var i Humaner
	//只要实现了此接口中所有方法的类型,那么这个类型的变量(接收类型)就可以给i赋值

	//实现接口
	s := &Student10{"liuhao", 666}

	i = s
	i.SayHi()

	t := &Teacher11{"bj", "go"}

	i = t
	i.SayHi()

	var str MyStr = "hello world"
	i = &str
	i.SayHi()

}
