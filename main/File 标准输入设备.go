package main

import (
	"fmt"
	"os"
)

func main() {

	//关闭后无法直接输出
	//os.Stdout.Close()
	//往标准输出设备(屏幕)写内容
	fmt.Println("are you ok")
	//标准设备文件(os.Stdout) 默认已经打开,用户可以直接使用
	os.Stdout.WriteString("are you ok? \n")

	//关闭后无法输入
	//os.Stdin.Close()  //关闭标准设备输入内容
	var a int
	fmt.Println("请输入a:")
	fmt.Scan(&a) //从标准输入设备中读取内容放在A中
	fmt.Println("a= ", a)

}
