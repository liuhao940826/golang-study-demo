package main

import "fmt"

type Person2 struct {
	name string
	sex  byte
	age  int
}

//这边的老师继承了person
type Teacher2 struct {
	Person2 //只有类型 没有名字 匿名字段继承了Person里面的成员
	id      int
	addr    string
	name    string
}

//数组长度在定义后无法再次修改 数组是值类型 每次传递产生一个副本  用slice来弥补不足
func main() {

	//推导
	t2 := Teacher2{Person2{"djy", 'm', 25}, 1, "shanghai", "liuhao"}
	fmt.Printf("t2 %+v \n", t2)
	//整体赋值
	t2.Person2 = Person2{"liuhao", 'm', 25}
	fmt.Printf("t2 %+v \n", t2)

	//问题在这里 相同的name 单个操作的时候怎么处理
	t2.name = "我是谁" //赋值的是自己的
	fmt.Printf("t2 %+v \n", t2)

	t2.Person2.name = "yoyo小姐姐"
	fmt.Printf("t2 %+v \n", t2)
}
