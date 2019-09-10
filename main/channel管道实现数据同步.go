package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string)
	//最后执行
	defer fmt.Println("主协程调用完毕.......")

	go func() {
		//子协程执行完毕了执行
		defer fmt.Println("子协程调用完毕.....")

		for i := 0; i < 2; i++ {

			fmt.Println("子协程 i=", i)
			time.Sleep(time.Second)
		}
		//循环结束往管道赋值 提供给主协程使用
		ch <- "我是子协程,调用完毕,通知主协程"
	}()

	str := <-ch //管道获取不到数据阻塞
	fmt.Println("输出str:", str)

}
