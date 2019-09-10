package main

import (
	"fmt"
	"time"
)

func main() {
	//每一秒 往ticker.C的这个通道中写入时间
	ticker := time.NewTicker(1 * time.Second)

	i := 0
	for {
		//从ticker.c的这个管道中 读内容 没有阻塞
		<-ticker.C

		i++
		fmt.Println("i=", i)

		if i == 5 {
			ticker.Stop()
			break
		}

	}

}
