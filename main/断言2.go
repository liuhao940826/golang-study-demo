package main

import (
	"fmt"
)

type Type2 struct {
	name string
	id   int
}

func main() {

	//创建一个切片
	i := make([]interface{}, 3)
	i[0] = 1
	i[1] = "hellow go"
	i[2] = Type1{"liuhao", 666}

	//类型断言
	//第一个返回值是下标 第二个返回值是对应的值
	for index, data := range i {

		switch value := data.(type) {

		case int:
			fmt.Printf("x[%d], 类型为int 值为%d \n", index, value)
		case string:
			fmt.Printf("x[%d], 类型string 值为%s \n", index, value)
		case Type2:
			fmt.Printf("x[%d], 类型Type 值为%v \n   ", index, value)
		}

	}

}
