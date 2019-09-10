package main

import (
	"fmt"
)

//select 监听channel
func main() {

	ch1 := make(chan int)

	ch2 := make(chan string)
	//select 的case 必须一个IO操作

	select {
	//当 1 2 都成功 的时候 无序的执行 1 2
	case <-ch1:
		fmt.Println("case1 的IO操作执行成功")
	case ch2 <- "1":
		fmt.Println("case2 的IO操作执行成功")
	default:
		fmt.Println("上面都没成功就是我.... 写我的话就变成轮训的了影响性能")

	}

}
