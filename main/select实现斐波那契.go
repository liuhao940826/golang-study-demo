package main

import (
	"fmt"
)

//select 监听channel
func main() {
	//保存数字
	ch1 := make(chan int) //数字通讯

	ch2 := make(chan bool) //程序是否结束
	//select 的case 必须一个IO操作

	//消费者从channel 读取内容
	go func() {
		for i := 0; i < 8; i++ {
			num := <-ch1
			fmt.Println(num)
		}
		//可以停止
		ch2 <- true

	}()

	//生产者从往channel 写入内容
	fibonacci(ch1, ch2)

}

//num 只写  bools 只读
func fibonacci(num chan<- int, bools <-chan bool) {

	x, y := 1, 1
	for {
		select {

		case num <- x:
			x, y = y, x+y
		case flag := <-bools:
			fmt.Println("flag", flag)
			return
		}
	}

}
