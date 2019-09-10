package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int) //创建一个无缓存的channel

	go func() {

		fmt.Println("子协程开始")

		for i := 0; i < 5; i++ {
			ch <- i //往chan写内容
			fmt.Printf("子协程 i= %d  \n", i)
		}
		//如果已经不需要写数据了就关闭channel
		close(ch)
		//关闭通道以后 无法在发送数据 直接报错  但是还是可以读的 但是如果nil是空的 读写都阻塞
		//ch<- 666
	}()

	time.Sleep(2 * time.Second)

	for {
		//如果为ok = true 管道没有关闭
		if num, ok := <-ch; ok == true {
			fmt.Print("num= \n", num)
		} else {
			//管道关闭 跳出循环
			break
		}
	}

}
