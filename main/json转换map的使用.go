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

	str := []string{"me", "you", "anyone"}
	//格式化 空字符变成了Tab键
	result, err := json.MarshalIndent(str, "", "	")

	if err != nil {
		fmt.Println("err", err)
		return
	}

	sliceStr := appendJSONSliceStr("新来的", string(result))
	str22 := appendJSONSliceStr("新来的Str", "我也是新来的")

	fmt.Println("新的sliceStr", sliceStr)
	fmt.Println("str22", str22)
}

func appendJSONSliceStr(str, jsonStr string) string {

	//只能判断已经有值了判断不了是slice还是string
	traceIds := &[]string{}
	//解析判断是否是数组
	err := json.Unmarshal([]byte(jsonStr), traceIds)
	if err == nil {
		//解析成功 原本已经是数组了,拼到后面去然后在抓json
		*traceIds = append(*traceIds, str)
		newTraceIds, _ := json.MarshalIndent(traceIds, "", "	")
		return string(newTraceIds)
	}

	//解析不成功,走到这边说明原本是个string 拼到数组里面去
	*traceIds = append(*traceIds, str, jsonStr)
	newTraceIds, _ := json.MarshalIndent(traceIds, "", "	")
	return string(newTraceIds)
}
