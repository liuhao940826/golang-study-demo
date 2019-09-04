package main

import (
	"fmt"
)

type Type1 struct {
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
		//第一个返回值 第二个返回判断结果的真假
		if value, ok := data.(int); ok == true {
			fmt.Printf("x[%d], 类型为int 值为%d \n", index, value)
		} else if value, ok := data.(string); ok == true {

			fmt.Printf("x[%d], 类型string 值为%s \n", index, value)
		} else if value, ok := data.(Type1); ok == true {
			fmt.Printf("x[%d], 类型string 值为%v \n   ", index, value)
		}

	}

}
