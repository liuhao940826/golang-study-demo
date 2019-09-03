package main

import (
	"fmt"
)

func main() {

	i := c()
	fmt.Print("值", i)

}

func c() (i int) {
	//这边的i 是获取到 return以后的值操作的 因为defer在这c函数的作用域中
	defer func() { i++ }()
	return 1
}

//协程批量插入
func batchDone() {
	func() {
		pos := "123456"
		defer fmt.Print("我最后打")
		for i := range pos {
			fmt.Println(i)
		}
	}()
}
