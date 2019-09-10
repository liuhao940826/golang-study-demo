package main

import (
	"fmt"
	"time"
)

//有容量的通道 在容量没满之前 是异步的 只有当容量满的情况下 会同步阻塞 取出以后才能写进去
func main() {

	//有缓存的channel 可以存储一定值的数据  使用有缓存的通道在goroutine之间同步数据
	//缓冲区满了不能够存东西 阻塞, 缓冲区空了不能够取东西阻塞
	//无缓冲区的是 存了就不取就阻塞
	// 容量就是3
	ch := make(chan int, 3)

	//chan的lens是缓冲区剩余个数 因为不带缓冲区就是0 ,cap(ch)是缓冲区大小
	fmt.Printf("len(ch) =%d cap(ch) =%d \n", len(ch), cap(ch))

	go func() {

		fmt.Println("子协程开始")

		for i := 0; i < 10; i++ {
			//主线程sleep了 容量只有3个 第一次放进来3个这边阻塞 0 1 2 然后等待主线程取出 然后才能在放入
			ch <- i //往chan写内容

			fmt.Printf("子协程 i= %d  len(ch) =%d,cap(ch) =%d\n", i, len(ch), cap(ch))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 10; i++ {
		num := <-ch
		fmt.Println("num=", num)
	}

}
