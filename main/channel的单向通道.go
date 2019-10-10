package main

import ()

//管道默认是双向的 既可以写也可以读     recv <- chan <- send
func main() {

	//无缓存的channel 会阻塞内不完成
	ch := make(chan int)

	//双向的隐式转向 单项channel
	var writeCh chan<- int = ch //只能写不能读

	writeCh <- 666 //写

	var readCh <-chan int = ch //只能读不能写

	<-readCh //读

	//这里会报错
	//<-writeCh 读
	//readCh<-666 写

	//单项无法转换成双向
	//var ch2 chan int =writeCh

}
