package main

import (
	"encoding/json"
	"fmt"
)

type JsonSlice struct {
	Company string `json:"company"` //转出来的时候取的名字
	Name    string `json:"name"`
	Age     int    `json:"-"`              //表示不进行序列化
	IsOk    bool   `json:"isReady,string"` // ,string表示序列化时候指定类型
}

func main() {

	strings := []string{"me", "you", "anyone"}
	//格式化 空字符变成了Tab键
	result, err := json.MarshalIndent(strings, "", "	")

	if err != nil {
		fmt.Println("err", err)
		return
	}

	fmt.Println("result", string(result))

	result2 := &[]string{}
	err := json.Unmarshal([]byte(result), result2)

	if err != nil {
		fmt.Println("err", err)
		return
	}

	fmt.Println("result2的内容", result2)

}
