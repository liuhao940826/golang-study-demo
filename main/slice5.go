package main

import "fmt"

//slice 不是真正意义上的动态数组 (切片不是数组或者数组指针) 他通过内部指针和相关属性引用数组片段,以实现变长方案

//slice 总是指向一个底层 array,slice的声明 也可以像array一样 只是不需要长度
//数组长度在定义后无法再次修改 数组是值类型 每次传递产生一个副本  用slice来弥补不足
func main() {
	//切片和底层数组的关系
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//截取 [loww,high,max] 取下标从low开始的元素 长度 =high-low  cap =max-low
	//新切片
	s1 := a[2:5]
	s1[1] = 666
	fmt.Printf("s1= \n", s1)
	fmt.Printf("a= \n", a)

	//另一个新切片
	s2 := s1[2:7]
	s2[2] = 777
	fmt.Printf("s2= \n", s2)
	fmt.Printf("a= \n", a)
}
