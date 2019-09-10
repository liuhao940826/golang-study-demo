package main

import (
	"fmt"
	"time"
)

func main() {

	//无缓存的channel 会阻塞内不完成
	ch := make(chan int, 0)

	//chan的lens是缓冲区剩余个数 因为不带缓冲区就是0 ,cap(ch)是缓冲区大小
	fmt.Printf("len(ch) =%d cap(ch) =%d \n", len(ch), cap(ch))

	go func() {

		fmt.Println("子协程开始")

		for i := 0; i < 4; i++ {
			fmt.Printf("子协程 i= %d \n", i)
			//放进来 但是下面不取 这边会阻塞
			ch <- i //往chan写内容
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++ {
		num := <-ch
		fmt.Println("num=", num)
	}

}
