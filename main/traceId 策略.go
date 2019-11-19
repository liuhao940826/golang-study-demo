package main

import (
	"encoding/json"
	"fmt"
	"github.com/jtolds/gls"
	"golang-study-demo/main/uuid"
)

const (
	TraceIdKey = "PM-TRACE-ID"
)

var (
	Mgr = gls.NewContextManager()
)

func main() {

	SetTraceIdSmart("aaa")

	id := GetTraceId()

	fmt.Println("内容", id)

	SetTraceIdSmart("bb")
}

func SetTraceId() {

	Mgr.SetValues(gls.Values{TraceIdKey: uuid.NewUuid()}, func() {})
}

func GetTraceId() string {
	if traceId, ok := Mgr.GetValue(TraceIdKey); ok {
		if traceId != nil {
			return traceId.(string)
		}
	}
	return ""
}

func SetTraceIdSmart(str string) {

	if str == "" {
		SetTraceId()
	} else {
		//如果设置的时候已经有了就把现有的放给父traceId 现有的放入traceId

		data, ok := Mgr.GetValue(TraceIdKey)

		if ok {
			if data != nil {
				if traceIdStr, ok := data.(string); ok == true {
					//只能判断已经有值了判断不了是slice还是string
					traceIds := &[]string{}
					//解析判断是否是数组
					err := json.Unmarshal([]byte(traceIdStr), traceIds)
					if err == nil {
						//解析成功 原本已经是数组了,拼到后面去然后在抓json
						*traceIds = append(*traceIds, str)
						newTraceIds, _ := json.MarshalIndent(traceIds, "", "	")
						Mgr.SetValues(gls.Values{TraceIdKey: newTraceIds}, func() {})
						return
					}

					//解析不成功,走到这边说明原本是个string 拼到数组里面去
					*traceIds = append(*traceIds, str, traceIdStr)
					newTraceIds, _ := json.MarshalIndent(traceIds, "", "	")
					Mgr.SetValues(gls.Values{TraceIdKey: newTraceIds}, func() {

						exist, b1 := Mgr.GetValue(TraceIdKey)

						fmt.Println("b1", b1)
						fmt.Println("exist", exist)

					})
				}
			}
		}

		//如果没有 就直接设置str
		Mgr.SetValues(gls.Values{TraceIdKey: str}, func() {

			exist, b1 := Mgr.GetValue(TraceIdKey)

			fmt.Println("b1", b1)
			fmt.Println("exist", exist)
		})

		value, b := Mgr.GetValue(TraceIdKey)

		fmt.Println("b", b)
		fmt.Println("value", value)
	}
}
