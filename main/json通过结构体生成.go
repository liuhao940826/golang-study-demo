package main

import (
	"encoding/json"
	"fmt"
)

type JsonStruct struct {
	Company string
	Name    string
	Age     int
}

type JsonStruct2 struct {
	Company string
	name    string
	age     int
}

func main() {

	jsonObj := JsonStruct{"哈哈哈", "liuhao", 25}

	buf, err := json.Marshal(jsonObj)

	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("buf", string(buf))

	fmt.Println("-------------------------------------")
	//如果结构体的字段类型小写的话 是不会打印出来的
	jsonObj2 := JsonStruct2{"哈哈哈", "liuhao", 25}

	buf2, err2 := json.Marshal(jsonObj2)

	if err != nil {
		fmt.Println("err", err2)
		return
	}
	fmt.Println("buf", string(buf2))

}
