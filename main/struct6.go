package main

import "fmt"

type Student5 struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

//结构体指针变量初始化
func main() {
	//结构体的相互赋值
	s := Student5{1, "s1", 'm', 18, "sh"}

	printStudent4(s) //值传递  形参无法更改实参
	fmt.Println("s", s)

	fmt.Println("==================")
	printStudent4Point(&s)
	fmt.Println("s", s)

}

func printStudent4Point(s *Student5) {

	(*s).name = "yoyo"
	fmt.Println("printStudent4Point", s)
}

func printStudent4(s Student5) {

	s.name = "yoyo"
	fmt.Println("printStudent4", s)
}
