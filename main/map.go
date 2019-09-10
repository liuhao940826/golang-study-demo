package main

import (
	"fmt"
)

func main() {

	//格式 map[key]value
	//key要求是唯一的 并且 key要支持== 和!=操作
	info := map[int]string{
		110: "liuhao",
		111: "hahah"}

	fmt.Printf("%v", info)

}
