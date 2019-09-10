package main

import (
	"fmt"
)

type Humaner4 interface { //子级接口东西少
	SayHi() //这个叫方法 只有声明 没有实现 由别的(自定义类型)实现
}

type Humaner5 interface { //超级接口 东西多
	Humaner4 //匿名字段继承 这个方法
	SayBye(bye string)
}

type Student13 struct {
	name string
	id   int
}

func (tmp *Student13) SayHi() {
	fmt.Printf("Student[%s ,%d] sayHi \n", tmp.name, tmp.id)
}

func (tmp *Student13) SayBye(bye string) {
	fmt.Println("跟你说拜拜", bye)
}

func main() {
	//超级接口可以转化为子级接口 反过来不可以
	var super Humaner5

	var son Humaner4

	super = &Student13{"liuhao", 666}

	son = super

	son.SayHi()

}
