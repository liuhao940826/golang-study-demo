package main

import "fmt"

//type的正确含义是起别名   这边给func类型的函数起了一个名字 叫funcType
type FuncType func(a, b int) int

//这边的name等同于 string类型 就是给string 起了一个别名
type name string

type stringFuncType func(str1, str2 string) string

//加
func add(a, b int) int {

	fmt.Println("I'm func1 ")

	return a + b
}

//减
func sub(a, b int) int {
	fmt.Println("I'm func2")
	return a - b
}

//打招呼
func sayHi(sayer, besayer string) string {
	return sayer + "对" + besayer + "说 hi"
}

//说告辞
func sayBye(sayer, besayer string) string {
	return sayer + "对" + besayer + "说 bye"
}

func main() {
	//方法1 的变量定义
	var intResult int
	var funcType1 FuncType
	//数据初始化
	a, b := 10, 20
	//函数类型赋值
	funcType1 = add
	//函数变量实际调用
	intResult = funcType1(a, b)
	fmt.Println("加法计算结果:", intResult)
	funcType1 = sub
	intResult = funcType1(a, b)
	fmt.Println("减法计算结果:", intResult)

	//方法2 的变量定义
	var stringResult string
	var funcType2 stringFuncType

	person1, person2 := "liuhao", "girls"
	funcType2 = sayHi
	stringResult = funcType2(person1, person2)
	fmt.Println("撩妹:", stringResult)
	funcType2 = sayBye
	stringResult = funcType2(person1, person2)
	fmt.Println("分手:", stringResult)

}
