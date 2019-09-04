package main

import "fmt"

//slice 不是真正意义上的动态数组 (切片不是数组或者数组指针) 他通过内部指针和相关属性引用数组片段,以实现变长方案

//slice 总是指向一个底层 array,slice的声明 也可以像array一样 只是不需要长度
//数组长度在定义后无法再次修改 数组是值类型 每次传递产生一个副本  用slice来弥补不足
func main() {
	//切片的截取

	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//截取 [loww,high,max] 取下标从low开始的元素 长度 =high-low  cap =max-low

	s1 := array[:] //[0:len(array):cap(array)]
	fmt.Printf("s1 = %v \n", s1)
	fmt.Printf("len=%d,cap=%d slice %v\n", len(s1), cap(s1), s1)

	data := array[0]
	fmt.Printf("data = %v \n", data)

	s2 := array[3:6:7]
	fmt.Printf("s2 = %v \n", s2)
	fmt.Printf("len=%d,cap=%d slice %v\n", len(s2), cap(s2), s2)

	//不写从low0开始  cap和high一样  常用
	s3 := array[:6]
	fmt.Printf("s3 = %v \n", s3)
	fmt.Printf("len=%d,cap=%d slice %v\n", len(s3), cap(s3), s3)

	//从3开始到结尾
	s4 := array[3:]
	fmt.Printf("s4 = %v \n", s4)
	fmt.Printf("len=%d,cap=%d slice %v\n", len(s4), cap(s4), s4)

}
