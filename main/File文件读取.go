package main

import (
	"fmt"
	"os"
)

func main() {

	path := "./demo.txt"

	readFile(path)

}

func readFile(path string) {
	//打开文件
	f, err := os.Open(path)

	if err != nil {
		return
	}

	defer f.Close()

	buf := make([]byte, 1024*2)
	//n代表从文件读取内容的长度
	n, err1 := f.Read(buf)

	if err1 != nil {
		return
	}

	fmt.Println("buf=", string(buf[:n]))

}
