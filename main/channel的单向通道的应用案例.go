package main

import "fmt"

//管道默认是双向的 既可以写也可以读     recv <- chan <- send
func main() {

	//创建一个双向通道
	ch := make(chan int)

	//生产者,生产数字,写入channel
	//这个是引用传递 就是地址本身
	go producer(ch)

	//消费者,从channel 读取内容 ,打印

	consumer(ch)

}

//这个只能读,不能写
func consumer(inner <-chan int) {
	//上面通道关闭以后这边range自动停止
	for num := range inner {
		fmt.Println("num=", num)
	}

}

//此通道只能写,不能读
func producer(out chan<- int) {

	for i := 0; i < 5; i++ {
		out <- i * i
	}
	close(out)
}
