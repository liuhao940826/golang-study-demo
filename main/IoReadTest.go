package main

import (
	"fmt"
	"os"
)

func main() {
	dir := "./../../templete/testTemplete.txt"

	path, _ := os.Getwd()

	fmt.Println("当前路径", path)

	fileObj, err := os.Open(dir)
	defer func() {
		if err := fileObj.Close(); err != nil {
			fmt.Println("异常")
		}
	}()

	if err != nil {
		fmt.Println("异常")
	}

}
