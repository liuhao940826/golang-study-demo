package main

import (
	"fmt"
)

//type FuncType func(num1,num2 int)  int

//加法
func Add(num1, num2 int) int {
	return num1 + num2
}

//减法
func Sub(num1, num2 int) int {
	return num1 - num2
}

//计算方法 行为同
func calc(num1, num2 int, selfFuncType FuncType) (result int) {
	fmt.Println("这是一个计算方法")
	result = selfFuncType(num1, num2)
	return

}

func main() {
	fmt.Println("callback")

	sum := calc(20, 10, Add)
	fmt.Println("结果:", sum)

	poor := calc(20, 10, Sub)
	fmt.Println("结果:", poor)

}
