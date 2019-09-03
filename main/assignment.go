package main

import "fmt"

//第 1 行，将 []int 声明为 IntSlice 类型。
type IntSlice []int

//为这个类型编写一个 Len 方法，提供切片的长度。
func (p IntSlice) Len() int { return len(p) }

//根据提供的 i、j 元素索引，获取元素后进行比较，返回比较结果。
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }

//根据提供的 i、j 元素索引，交换两个元素的值。
func (p IntSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

//上面的先不管  使用 Go 语言进行排序时就需要使用交换，代码如下：
func main() {
	//多重赋值时，变量的左值和右值按从左到右的顺序赋值。

	//多重赋值在 Go 语言的错误处理和函数返回值中会大量地使用。
	var a int = 100
	var b int = 200
	b, a = a, b
	fmt.Println(a, b)

}
