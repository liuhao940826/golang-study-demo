package main

import (
	"fmt"
	"runtime"
)

//main 方法 叫做主协程
func main() {

	//子协程
	go func() {

		for i := 0; i < 5; i++ {

			fmt.Println("go")

		}
	}()

	for i := 0; i < 2; i++ {

		//主线程让出cpu的时间片 让出时间片不是绝对的
		runtime.Gosched()

		fmt.Println("hello")

	}

}
