package main

import (
	"fmt"
)

//slice 不是真正意义上的动态数组 (切片不是数组或者数组指针) 他通过内部指针和相关属性引用数组片段,以实现变长方案

//slice 总是指向一个底层 array,slice的声明 也可以像array一样 只是不需要长度
//数组长度在定义后无法再次修改 数组是值类型 每次传递产生一个副本  用slice来弥补不足
func main() {

	//切片的创建方式  自动推导  make函数(类型 ,长度,容量)

	//自动推导类型,同时初始化
	//s1 := []int{1, 2, 3, 4}
	//fmt.Println(s1)

	//借助make , 格式make(切片类型, 长度,cap)
	s2 := make([]int, 5, 10)
	fmt.Printf("len=%d,cap=%d slice %v\n", len(s2), cap(s2), s2)
	//不指定容量 容量和长度一样
	s3 := make([]int, 5)
	fmt.Printf("len=%d,cap=%d slice %v\n", len(s3), cap(s3), s3)

}
