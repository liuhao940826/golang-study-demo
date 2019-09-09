package main

import "fmt"

type TestOut struct {
	Id        int
	TestInner []struct {
		name string
	}
}

//数组长度在定义后无法再次修改 数组是值类型 每次传递产生一个副本  用slice来弥补不足
func main() {
	var resp TestOut
	//append 在切片末尾追加数字 如果不够就扩容
	TestInner := resp.TestInner
	fmt.Println("first", TestInner)
	resp.TestInner = append(TestInner, struct{ name string }{name: "liuhao"})
	fmt.Println("end", resp.TestInner)

}
