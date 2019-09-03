package main

import "fmt"

func main() {

	//var 变量名 变量类型  标准格式
	var a int = 10
	fmt.Println(a)

	//批量声明
	var (
		b int    = 20
		c string = "hello world"
	)
	fmt.Println(b, c)

	//简短格式  名字 := 表达式
	//因为简洁和灵活的特点，简短变量声明被广泛用于大部分的局部变量的声明和初始化。
	// var 形式的声明语句往往是用于需要显式指定变量类型地方，或者因为变量稍后会被重新赋值而初始值无关紧要的地方。
	x := 100
	y, z := 1, "abc"
	fmt.Println("简短格式:", x, y, z)

}
