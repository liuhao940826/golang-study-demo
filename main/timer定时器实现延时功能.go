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
	<-timer.C //没有数据前会阻塞
	fmt.Println("我是2秒后要说的话..........")

}

func sleep() {
	time.Sleep(2 * time.Second)
	fmt.Println("我也延时2秒")
}

func afterTime() {
	//本身返回一个管道 定时2秒 阻塞2秒 2秒后产生一个事件 往channel写内容
	<-time.After(2 * time.Second)
	fmt.Println("谁不是2秒真男人")
}
