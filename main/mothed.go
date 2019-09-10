package main

import (
	"fmt"
)

//面向过程
func Add01(a, b int) int {
	return a + b
}

type sInt int

//tmp 叫做接受者 接受者就是一个参数
func (tmp sInt) Add02(other sInt) sInt {
	return tmp + other
}

func main() {

	var result int
	result = Add01(1, 1) //普通函数调用方式
	fmt.Println("result", result)

	//面向方法的调用
	var a sInt = 2
	//调用方法格式:  变量名.函数(所需参数)
	result2 := a.Add02(3)
	fmt.Println("result2", result2)

}
