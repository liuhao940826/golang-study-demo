package main

import (
	"encoding/json"
	"fmt"
)

type JsonMap struct {
	Company string `json:"company"` //转出来的时候取的名字
	Name    string `json:"name"`
	Age     int    `json:"-"`              //表示不进行序列化
	IsOk    bool   `json:"isReady,string"` // ,string表示序列化时候指定类型
}

func main() {

	m := make(map[string]interface{})

	m["company"] = "新型公司"
	m["age"] = 1
	m["member"] = []string{"me", "you", "anyone"}
	//格式化 空字符变成了Tab键
	result, err := json.MarshalIndent(m, "", "	")

	if err != nil {
		fmt.Println("err", err)
		return
	}

	fmt.Println("result", string(result))

	fmt.Println("result", string(result))

}
