package main

import (
	"fmt"
)

type Humaner2 interface {
	SayHi() //这个叫方法 只有声明 没有实现 由别的(自定义类型)实现
}

type Humaner3 interface {
	Humaner2 //匿名字段继承 这个方法
	SayBye(bye string)
}

type Student12 struct {
	name string
	id   int
}

func (tmp *Student12) SayHi() {
	fmt.Printf("Student[%s ,%d] sayHi \n", tmp.name, tmp.id)
}

func (tmp *Student12) SayBye(bye string) {
	fmt.Println("跟你说拜拜", bye)
}

func main() {
	//接口的继承
	var i Humaner3

	s := &Student12{"liuhao", 6}

	i = s

	i.SayHi()
	i.SayBye("you")

}
