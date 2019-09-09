package main

import (
	"fmt"
	"strconv"
)

func appendS(s *[]int, num int) {
	//拿到地址的值
	*s = append(*s, num)
	fmt.Printf("里面的 len=%d,cap=%d slice %v\n", len(*s), cap(*s), *s)
}

//数组长度在定义后无法再次修改 数组是值类型 每次传递产生一个副本  用slice来弥补不足
func main() {

	//append 在切片末尾追加数字 如果不够就扩容
	s1 := []int{}

	fmt.Printf("len=%d,cap=%d slice %v\n", len(s1), cap(s1), s1)

	//在元切片的末尾添加元素
	s1 = append(s1, 1)
	s1 = append(s1, 2)
	s1 = append(s1, 3)
	fmt.Printf("len=%d,cap=%d slice %v\n", len(s1), cap(s1), s1)

	appendS(&s1, 5)
	appendS(&s1, 5)
	appendS(&s1, 5)
	fmt.Printf("len=%d,cap=%d slice %v\n", len(s1), cap(s1), s1)
	//s2 := []int{1, 2, 3}
	//fmt.Printf("len=%d,cap=%d slice %v\n", len(s2), cap(s2), s2)
	//
	//s2 = append(s2, 4)
	//s2 = append(s2, 5)
	//s2 = append(s2, 6)
	//fmt.Printf("len=%d,cap=%d slice %v\n", len(s2), cap(s2), s2)

	//第二个参数代表十进制
	formatInt := strconv.FormatInt(10, 10)

	fmt.Println("fmtInt", formatInt)

}
