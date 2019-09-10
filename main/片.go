package main

import (
	"fmt"
)

//数组长度在定义后无法再次修改 数组是值类型 每次传递产生一个副本  用slice来弥补不足
func main() {

	//数组  []中的元素个数必须是常量
	array := [5]int{1, 2, 3}
	fmt.Println(array)

	//slice 不是真正意义上的动态数组 (切片不是数组或者数组指针) 他通过内部指针和相关属性引用数组片段,以实现变长方案

	//slice 总是指向一个底层 array,slice的声明 也可以像array一样 只是不需要长度
	//切片
	slice := array[0:3:5]
	fmt.Println(slice)

	//[low :high:max]
	//low :下标的起点  high:下标的终点 左闭右开 len =high -low(另一种记忆方式) cap =max -low 容量
	slice2 := array[1:4:5]
	fmt.Println(slice2)

}
