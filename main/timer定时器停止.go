package main

import (
	"fmt"
	"time"
)

//time 是时间到了触发一次
func main() {
	timer := time.NewTimer(3 * time.Second)

	go func() {
		<-timer.C
		fmt.Println("子协程 3秒到了.......")
	}()

	//这边停了子协程里面的 timer.C无法执行了 就不出来了
	timer.Stop()

	for {
		fmt.Println("不让你结束")
	}

}
