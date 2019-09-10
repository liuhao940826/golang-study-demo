package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

//这边的老师继承了person
type Teacher struct {
	Person //只有类型 没有名字 匿名字段继承了Person里面的成员
	id     int
	addr   string
}

//数组长度在定义后无法再次修改 数组是值类型 每次传递产生一个副本  用slice来弥补不足
func main() {

	var t1 Teacher = Teacher{Person{"liuhao", 'm', 25}, 1, "shanghai"}

	fmt.Println(t1)

	//推导
	t2 := Teacher{Person{"djy", 'm', 25}, 1, "shanghai"}

	fmt.Printf("t2 %+v \n", t2)

	//指定成员初始化
	t3 := Teacher{id: 1}
	fmt.Printf("t3 %+v \n", t3)

	//指定成员初始化
	t4 := Teacher{Person: Person{name: "jubu"}, id: 1}
	fmt.Printf("t4 %+v \n", t4)

}
