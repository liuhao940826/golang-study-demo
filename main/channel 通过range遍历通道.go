package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("我是测试拼接" + strconv.Itoa(666))

	ch := make(chan string)
	//最后执行
	defer fmt.Println("主协程调用完毕.......")

	go func() {
		//子协程执行完毕了执行
		defer fmt.Println("子协程调用完毕.....")

		for i := 0; i < 4; i++ {

			//循环结束往管道赋值 提供给主协程使用
			ch <- "我是值" + strconv.Itoa(i)
		}
		//不需要数据的时候关闭  关闭后无法在塞入数据
		close(ch)

	}()
	//关闭的时候自动跳出这个循环不读了
	for str := range ch {
		fmt.Println("管道的数据:", str)
	}

}
