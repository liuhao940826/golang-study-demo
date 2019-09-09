package main

import (
	"encoding/json"
	"fmt"
)

type JsonTag struct {
	Company string `json:"company"` //转出来的时候取的名字
	Name    string `json:"name"`
	Age     int    `json:"-"`              //表示不进行序列化
	IsOk    bool   `json:"isReady,string"` // ,string表示序列化时候指定类型
}

func main() {

	jsonObj := JsonTag{"哈哈哈", "liuhao", 25, true}

	buf, err := json.Marshal(jsonObj)

	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("buf", string(buf))

}
