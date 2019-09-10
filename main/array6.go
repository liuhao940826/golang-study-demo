package main

import (
	"fmt"
)

//p指向实现数组a 他是指向数组 他是数组指针
//*p 代表指针所指向的内存,就是实参a
func modifyP(p *[5]int) {
	//指针指向 p的那个地址  取第0个元素
	(*p)[0] = 666
	fmt.Println("modify", *p)
}

func main() {

	//数组比较   只支持== 和!=比较  比较是不是每个元素都一样 2个数组比较 数组类型要一样 [n]int 为数组类型
	var a = [5]int{1, 2, 3, 4, 5}
	modifyP(&a)
	//值传递外层不会改
	fmt.Println("main", a)
}
