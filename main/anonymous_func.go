package main

import "fmt"

func main() {

	a, b := 10, 20

	f1 := func() {
		fmt.Println("测试匿名函数", a, b)
	}
	//调用匿名函数
	f1()
	//如果匿名函数没有被使用就定义 需要在最后的()代表,同时调用这个函数 因为golang中写的东西一定要被用
	func() {
		fmt.Println("没有使用的测试匿名函数", a, b)
	}()
	//给无参无返回的func起一个别名  这个只是定义一个类型 这个是类型
	type Anonymous1 func()
	//定义一个 Anonymous1类型的变量 实际就是 func()类型的
	var f2 Anonymous1
	f2 = f1
	//调用 别名方法变量
	f2()

}
