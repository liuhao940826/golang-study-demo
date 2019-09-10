package main

import (
	"fmt"
	"time"
)

//time 是时间到了触发一次
func main() {
	//创建一个定时器 设置时间为2秒  2s后 往time 通道写内容(当前时间)
	timer := time.NewTimer(2 * time.Second)

	fmt.Println("当前时间", time.Now())

	// 2s后,往timer.C写数据,有数据后可以读取
	t := <-timer.C //没有数据前会阻塞
	fmt.Println("t:", t)

	//验证timer时间到了 只会响应一次
	timer2 := time.NewTimer(1 * time.Second)

	for {
		<-timer2.C
		fmt.Println("时间到")
		//这里会死锁
	}

}
