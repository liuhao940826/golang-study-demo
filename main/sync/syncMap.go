package main

import (
	"fmt"
	"sync"
)

//Store 写入
//Load 读取，返回值有两个，第一个是value，第二个是bool变量表示key是否存在
//Delete 删除
//LoadOrStore 存在就读，不存在就写
//Range 遍历，注意遍历的快照
func main() {

	sMap := sync.Map{}

	sMap.Store("test", "1")

	value1, _ := sMap.Load("test")

	fmt.Println("value1:", value1)

	actual, loaded := sMap.LoadOrStore("test", "2")

	fmt.Printf("actual ====%s  loaded ====%v \n", actual, loaded)

	//加载test 拿到的是1
	value2, _ := sMap.Load("test")

	fmt.Println("value2:", value2)

}
