package main

import "fmt"

func main() {

	//map 做函数参数 引用传递
	//map的常用 操作 遍历和判断key 是否存在
	m := map[int]string{1: "mike", 2: "yoyo", 3: "c++"}
	testMap(m)

	fmt.Println(m)
}

func testMap(m map[int]string) {
	delete(m, 1) //删除key为1的内容
}
