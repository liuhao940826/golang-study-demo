package main

import (
	"fmt"
	"reflect"
)

func main() {
	var circle float64 = 6.28

	var icir interface{}
	icir = circle

	valueref := reflect.ValueOf(icir)

	t := reflect.TypeOf(icir)
	fmt.Println("type的类型", t)

	fmt.Println("这个是valueof", valueref)
	fmt.Println(valueref.Interface())

	//value转type
	i := valueref.Type()
	fmt.Println("是否相等 true", t == i)
	//获取他对应的值
	y := valueref.Interface()
	fmt.Println("v.interface():", y)
}
