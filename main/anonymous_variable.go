package main

import "fmt"

//匿名变量
func GetData() (int, int) {
	return 100, 200
}

func main() {
	//***匿名变量不占用命名空间，不会分配内存。匿名变量与匿名变量之间也不会因为多次声明而无法使用。

	//只需要获取第一个返回值，所以将第二个返回值的变量设为下画线。
	a, _ := GetData()
	//将第一个返回值的变量设为匿名。
	_, b := GetData()

	fmt.Println(a, b)
}
