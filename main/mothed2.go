package main

import (
	"fmt"
)

type Person5 struct {
	name string
	sex  byte
	age  int
}

//这边的老师继承了person
type Teacher5 struct {
	*Person5 // 结构体匿名字段 只有类型 没有名字 匿名字段继承了Person里面的成员
	int      //基础匿名字段
}

func (tmp Person5) PrintInfo() {
	fmt.Println("tmp=", tmp)
}

//通过一个函数,给成员赋值
func (tmp *Person5) SetInfo(name string, sex byte, age int) {

	tmp.name = name
	tmp.age = age
	tmp.sex = sex
}

func main() {
	//定义同事初始化
	p := Person5{"mike", 'm', 25}
	p.PrintInfo()
	fmt.Println("===========================")
	//定义一个变量

	var p2 Person5
	(&p2).SetInfo("yoyo", 'f', 22)
	p2.PrintInfo()

}
