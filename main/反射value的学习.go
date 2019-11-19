package main

import (
	"fmt"
	"reflect"
)

type reflectUserType struct {
	name string
	id   int
}

type selfInterface interface {
}

func main() {

	r := reflectUserType{name: "alan", id: 1}

	v := reflect.ValueOf(r)
	//结果直接是 alan 1 这个结构题里面所有的值
	fmt.Println("valueOf 返回的内容", v)
	//返回V持有的
	fmt.Println("value 的kind", v.Kind())

	fmt.Println("value 的type", v.Type())
	//不是interface 类型的会报错的
	//fmt.Println("value 的elem",v.Elem())
}
