package main

import "fmt"

type mystr string //自定义类型 给一个类型改名

type Person3 struct {
	name string
	sex  byte
	age  int
}

//这边的老师继承了person
type Teacher3 struct {
	Person3 // 结构体匿名字段 只有类型 没有名字 匿名字段继承了Person里面的成员
	int     //基础匿名字段
	mystr   //自定义的匿名字段 就是string

}

//数组长度在定义后无法再次修改 数组是值类型 每次传递产生一个副本  用slice来弥补不足
func main() {

	//推导
	t2 := Teacher3{Person3{"djy", 'm', 25}, 1, "liuhao"}
	fmt.Printf("t2 %+v \n", t2)
	t2.int = 777
	fmt.Printf("t2 %+v \n", t2)
}
