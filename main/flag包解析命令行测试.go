package main

import (
	"flag"
	"fmt"
)

var constIp int

func main() {

	//名字: 我的ip名字  值 1234 提示语 help message for flagname
	var ip = flag.Int("我的ip名字", 1234, "help message for flagname")

	fmt.Println(*ip)
	//另一种赋值方式
	flag.IntVar(&constIp, "other ip", 8888, "other ip help message ")

	fmt.Println("另一个ip", constIp)

}
