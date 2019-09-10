package main

import (
	"container/list"
	"fmt"
)

type pointerStudy struct {
	id int
}

type pointerTypeStudy struct {
	id *int
}

type Upd map[string]interface{}

func testBool(flag *bool) {
	*flag = true
}

//看第一行
// *操作符作为右值时，意义是取指针的值；作为左值时，也就是放在赋值操作符的左边时，表示 a 指向的变量。

// 其实归纳起来，*操作符的根本意义就是操作指针指向的变量。当操作在右值时，就是取指向变量的值；当操作在左值时，就是将值设置给指向的变量。
func main() {
	//0xc000090000 指针变量类型 *int64  暂时理解为类型*(指针) 指向的空间类型为int64 i是那片空间的地址
	i := new(int64)
	fmt.Printf("new的结果: %v i的类型%T \n", i, i)

	i2 := &i
	fmt.Printf("i这个内存地址的地址 %p 这个内存地址的值等于i的内存地址 %v \n", i2, *i2)

	fmt.Println("*i2 == i 指针i2的值等于 i指针的值", *i2 == i)

	fmt.Printf("new的值 %v \n", **i2)
	//
	//var a=10;
	//p := &a
	//
	//fmt.Println("p的值",*p)

	study := pointerStudy{id: 1}

	id := study.id

	fmt.Println("id", id)
	//id2是和id3一样的
	id2 := &(study.id)
	fmt.Println("id的地址", id2)
	id3 := &study.id
	fmt.Println("id的地址", id3)

	//获取study的指针
	i3 := &study
	fmt.Printf("study的地址 %p 类型 %T \n", i3, i3)
	//通过地址获取 i3指向的变量study
	(*i3).id = 10

	fmt.Println("study", study)

	//studiesSlice := make([]pointerStudy, 5)
	pointId := 1

	typeStudy := pointerTypeStudy{id: &pointId}

	typeId := *typeStudy.id

	fmt.Println("typeId=", typeId)

	flag := false
	fmt.Println("flag=", flag)
	testBool(&flag)

	fmt.Println("flag=", flag)

	noticeUserList := list.New()
	fmt.Printf("noticeUserList= %v \n", noticeUserList)
	fmt.Printf("noticeUserList Front= %v\n", noticeUserList.Front())
	pushBakc(noticeUserList)
	fmt.Printf("noticeUserList= %v\n", noticeUserList)
	fmt.Printf("noticeUserList Front= %v\n", noticeUserList.Front())

	for e := noticeUserList.Front(); e != nil; e = e.Next() {
		fmt.Println("value", e.Value)
	}
}

func pushBakc(noticeUserList *list.List) {
	//(*noticeUserList).PushBack(1)

	noticeUserList.PushBack(1)
}

func putMap() {

}
