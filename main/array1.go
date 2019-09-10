package main

import (
	"fmt"
)

func main() {

	//数组  []中的元素个数必须是常量
	var ids [50]int

	for i := 0; i < len(ids); i++ {
		ids[i] = i + 1
		fmt.Printf("id[%d] =%d \n", i, ids[i])

	}

}
