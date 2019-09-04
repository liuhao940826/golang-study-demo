package main

import "fmt"

//数组长度在定义后无法再次修改 数组是值类型 每次传递产生一个副本  用slice来弥补不足
func main() {

	//append 在切片末尾追加数字 如果不够就扩容  如果超过原来的容量通常以3倍扩容
	s := make([]int, 0, 1)

	oldCap := cap(s)
	for i := 0; i < 8; i++ {
		s = append(s, i)
		if newCap := cap(s); oldCap < newCap {
			fmt.Printf("cap %d===>%d \n ", oldCap, newCap)
			oldCap = newCap
		}

	}

}
