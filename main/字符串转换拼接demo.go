package main

import (
	"fmt"
	"strings"
)

func main() {

	//secondInterval := 600
	//
	//timeStr := fmt.Sprintf("%s"+strconv.Itoa(secondInterval)+"%s", "-", IntervalSecondUnit)
	//
	//fmt.Println("转化你的结果:", timeStr)

	//str:= "http://www.baidu.com"
	//
	//index := strings.Index(str, "?")
	//
	//fmt.Println("位置下标",index)
	//
	//new := str[:strings.Index(str, "?")]
	//
	//fmt.Println("新的字符串",new)

	url := "http://www.baidu.com?123"
	index := strings.Index(url, "?")

	if index != -1 {
		url = url[:index]
	}

	fmt.Println("新的字符串", url)

}
