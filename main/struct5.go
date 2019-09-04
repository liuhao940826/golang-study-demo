package main

import "fmt"

type Student4 struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

//结构体指针变量初始化
func main() {
	//结构体的相互赋值
	s1 := Student4{1, "s1", 'm', 18, "sh"}
	s2 := Student4{1, "s1", 'm', 18, "sh"}
	s3 := Student4{1, "s1", 'm', 18, "1"}

	fmt.Println("s1=====s2", s1 == s2)
	fmt.Println("s2=====s3", s3 == s2)
	fmt.Println("s1=====s3", s1 == s3)
	var tmp Student4

	tmp = s3
	fmt.Println("tmp", tmp)

}
