package main

import "fmt"

type mystr2 string

type Person4 struct {
	name string
	sex  byte
	age  int
}

//这边的老师继承了person
type Teacher4 struct {
	*Person4 // 结构体匿名字段 只有类型 没有名字 匿名字段继承了Person里面的成员
	int      //基础匿名字段
	mystr2   //自定义的匿名字段 就是string

}

//结构体指针类型匿名字段
func main() {

	//推导 初始化指针类型
	t2 := Teacher4{&Person4{"djy", 'm', 25}, 1, "liuhao"}
	fmt.Printf("t2 %+v \n", t2)

	var t1 Teacher4
	//ti.person4这个时候是一个地址
	t1.Person4 = new(Person4)
	t1.name = "新的指针地址"
	fmt.Printf("t1 %+v \n", t1)

}
