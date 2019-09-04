package main

import "fmt"

type Person1 struct {
	name string
	sex  byte
	age  int
}

//这边的老师继承了person
type Teacher1 struct {
	Person1 //只有类型 没有名字 匿名字段继承了Person里面的成员
	id      int
	addr    string
}

//数组长度在定义后无法再次修改 数组是值类型 每次传递产生一个副本  用slice来弥补不足
func main() {

	//推导
	t2 := Teacher1{Person1{"djy", 'm', 25}, 1, "shanghai"}
	//也可以单个赋值
	t2.name = "aaa"
	t2.sex = 'w'
	fmt.Printf("t2 %+v \n", t2)
	//整体赋值
	t2.Person1 = Person1{"liuhao", 'm', 25}
	fmt.Printf("t2 %+v \n", t2)

}
