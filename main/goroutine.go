package main

import (
	"fmt"
	"time"
)

func goRountineFunc() {

	for {
		fmt.Println("this is a  newTask")
		time.Sleep(time.Second)
	}

}

//main 方法 叫做主协程
func main() {

	//开线程去处理这个东西
	go goRountineFunc()

	//如果主协程结束 子协程也结束
	i := 0
	for {
		i++

		fmt.Println("this is a  newTask")
		time.Sleep(time.Second)

		if i == 2 {
			break
		}

	}

}
