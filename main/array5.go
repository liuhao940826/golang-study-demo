package main

import (
	"fmt"
)

//这个是值传递
func modify(a [5]int) {
	a[0] = 666
	fmt.Println("modify", a)
}

func main() {

	//数组比较   只支持== 和!=比较  比较是不是每个元素都一样 2个数组比较 数组类型要一样 [n]int 为数组类型
	var a = [5]int{1, 2, 3, 4, 5}
	modify(a)
	//值传递外层不会改
	fmt.Println("main", a)
}
