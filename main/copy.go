package main

import "fmt"

//数组长度在定义后无法再次修改 数组是值类型 每次传递产生一个副本  用slice来弥补不足
func main() {

	//append 在切片末尾追加数字 如果不够就扩容  如果超过原来的容量通常以3倍扩容

	orgSlice := []int{1, 2}
	dstSlice := []int{6, 6, 6, 6, 6, 6}

	copy(dstSlice, orgSlice)
	fmt.Printf("%v", dstSlice)
	//{1,2,6,6,6,6} 把原来 1,2 替换到 dst的第一第二个

}
