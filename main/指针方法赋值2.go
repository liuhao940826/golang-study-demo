package main

import "fmt"

//*操作符作为右值时，意义是取指针的值；作为左值时，也就是放在赋值操作符的左边时，表示 a 指向的变量。
// 其实归纳起来，*操作符的根本意义就是操作指针指向的变量。当操作在右值时，就是取指向变量的值；当操作在左值时，就是将值设置给指向的变量。

type ponit2 struct {
	id   *int64
	name *string
}

type ponit3 struct {
	id   int64
	name string
}

type ponit4 struct {
	id   []int64
	name string
}

type ponit5 struct {
	id   int64
	name string
}

//数组长度在定义后无法再次修改 数组是值类型 每次传递产生一个副本  用slice来弥补不足
func main() {

	varint64 := int64(1)

	name := "liuhao"

	entity := &ponit2{
		id:   &varint64,
		name: &name,
	}

	fmt.Printf("entity2 id= %d name=%v \n ", *entity.id, *entity.name)
	testPonit2(entity)

	fmt.Printf("entity2 id= %d name=%v \n ", *entity.id, *entity.name)

	entity3 := &ponit3{
		id:   666,
		name: "hahah",
	}

	fmt.Printf("entity3 id= %d name=%v \n ", entity3.id, entity3.name)

	testPonit3(entity3)

	fmt.Printf("entity3 id= %d name=%v \n ", entity3.id, entity3.name)

	entity5 := &ponit5{
		id:   666,
		name: "hahah",
	}

	fmt.Printf("entity5 id= %d name=%v \n ", entity5.id, entity5.name)

	testPonit5(entity5)

	fmt.Printf("entity5 id= %d name=%v \n ", entity5.id, entity5.name)

	testPonit5_2(entity5)

	fmt.Printf("entity5_2 id= %d name=%v \n ", entity5.id, entity5.name)

}

func testPonit3(entity *ponit3) {

	entity.id = int64(11)

	entity.name = "nimama"

}

func testPonit2(entity *ponit2) {

	oint64 := int64(2)

	name := "nibaba"

	entity.id = &oint64

	entity.name = &name

}

//指针变量.id和   (*指针变量).id一样的
func testPonit5(entity *ponit5) {

	entity.id = int64(11)

	entity.name = "dsadsa"

}

func testPonit5_2(entity *ponit5) {

	(*entity).id = int64(111111)

	(*entity).id = int64(1111)

	(*entity).name = "qweqwe"

	var a int = 10

}
