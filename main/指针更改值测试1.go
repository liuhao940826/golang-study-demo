package main

import "fmt"

//*操作符作为右值时，意义是取指针的值；作为左值时，也就是放在赋值操作符的左边时，表示 a 指向的变量。
// 其实归纳起来，*操作符的根本意义就是操作指针指向的变量。当操作在右值时，就是取指向变量的值；当操作在左值时，就是将值设置给指向的变量。
func test1(memberFilter *bool, issueRelationIds *[]int64) {
	*memberFilter = true
	//新数组
	//int64s := []int64{1, 2, 3, 4}
	//issues:=&int64s

	for _, v := range *issueRelationIds {

		fmt.Println("片中内容: %+v", v)

	}
}

//数组长度在定义后无法再次修改 数组是值类型 每次传递产生一个副本  用slice来弥补不足
func main() {
	memberFilter := false
	issueRelationIds := &[]int64{1, 2, 3, 4}

	fmt.Println("issueRelationIds的长度%d 内容%v", len(*issueRelationIds), issueRelationIds)
	test1(&memberFilter, issueRelationIds)

	fmt.Printf("issueRelationIds的长度%d 内容%v", len(*issueRelationIds), issueRelationIds)

	//
	//fmt.Println("memberFilter",memberFilter)
	//
	//issueRelationPosInterface := make([]interface{}, 0, 10)
	//
	//fmt.Printf("类型  %T",issueRelationPosInterface)

}
