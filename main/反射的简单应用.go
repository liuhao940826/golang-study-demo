package main

import (
	"fmt"
	"reflect"
)

type reflectUser struct {
	Id   int64
	Name string
	Age  int64
}

func (u reflectUser) Helllo(name string) {
	fmt.Printf(" %s  Hellow world......................", name)
}

func Info(i interface{}) {
	//获取接口类型
	t := reflect.TypeOf(i)

}
