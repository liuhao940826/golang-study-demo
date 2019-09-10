package main

import (
	"fmt"
)

func main() {

	//map的常用 操作 遍历和判断key 是否存在
	m4 := map[int]string{1: "mike", 2: "yoyo", 3: "c++"}

	fmt.Println("m4=", m4)

	for key, value := range m4 {
		fmt.Printf("%d=======>%s\n", key, value)
	}

	//如何判断一个key是否存在
	//第一个返回值为key所对应的value 第二个返回值为key 是否存在的条件 存在 ok 为true

	//map直接取值可以获取一个值和一个是否存在
	value, ok := m4[1]

	if ok == true {
		fmt.Println("m[1]=", value)
	} else {
		fmt.Println("key不存在")
	}

}
