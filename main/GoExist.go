package main

import (
	"fmt"
	"runtime"
)

func test() {
	//这个会在当前方法内最后执行
	defer fmt.Println("ccccccccccccc")

	//终止所在协程  协程里面bbbb没有办法去执行了 defer 这里 还没执行 结果是 aaaa cccc
	runtime.Goexit()

	fmt.Println("ddddddddddddddddd")
}

func main() {

	go func() {
		fmt.Println("aaaaaaaaaaaaaaaaa")
		test()
		fmt.Println("bbbbbbbbbbbbbbb")
	}()

	for {
		i := 0
		i++

	}
}
