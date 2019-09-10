package main

import (
	"fmt"
)

func main() {
	//第二个参数代表长度 第三个代表容量
	ints := make([]int, 10, 10)
	fmt.Println(ints)
	fmt.Printf("slice的长度 %d,slice的容量 %d ,slice的内容 %v", len(ints), cap(ints), ints)

}
