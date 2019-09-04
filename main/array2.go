package main

import (
	"fmt"
)

func main() {

	//数组  []中的元素个数必须是常量  [10] int 整体看做一个类型 [5] int 整体看做一个类型
	var a [10]int
	var b [5]int

	fmt.Printf("len(a) =%d len(b)=%d \n", len(a), len(b))

	//[]里面必须写常量不可以用变量代替
	//n:=10
	//var c[n]int

	//这里的0 叫做下标 这个时候可以是 变量或者常量
	a[0] = 1

	for i := 0; i < len(a); i++ {
		a[i] = i + 1
	}

	m := 0

	fmt.Printf("a[%d]常量中的数字%d a[%d]变量下标中的数字%d \n", 0, a[0], m, a[m])

	for i, data := range a {
		fmt.Printf("a[%d] =%d \n", i, data)
	}

}
