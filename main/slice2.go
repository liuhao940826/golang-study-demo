package main

import (
	"fmt"
)

//数组长度在定义后无法再次修改 数组是值类型 每次传递产生一个副本  用slice来弥补不足
func main() {

	//数组[]中的元素个数必须是常量,不能够修改长度
	array := [5]int{}
	fmt.Printf("len=%d,cap=%d \n", len(array), cap(array))
	//slice 切片 []里面为空 或者为... 切片的长度或容量可以不固定
	s := []int{}
	fmt.Printf("len=%d,cap=%d \n", len(s), cap(s), s)

	s = append(s, 11) //给切片末尾追加一个成员
	fmt.Printf("len=%d,cap=%d \n", len(s), cap(s), s)

}
