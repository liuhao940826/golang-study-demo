package main

import (
	"fmt"
	"time"
)

//time 是时间到了触发一次
func main() {
	timer := time.NewTimer(3 * time.Second)

	ok := timer.Reset(1 * time.Second)

	fmt.Println("ok:=", ok)

	<-timer.C
	fmt.Println("时间到")

}
