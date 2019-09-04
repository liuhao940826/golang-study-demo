package main

import "fmt"

func main() {
	//全部初始化
	var a [5]int = [5]int{1, 2, 3, 4, 5}

	fmt.Println("a=", a)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b=", b)
	//部分初始化   其他的默认0
	c := [5]int{1, 2, 3}
	fmt.Println("c=", c)

	//指定初始化 下标2的为10 下标4的为50  其他的默认0
	d := [5]int{2: 10, 4: 50}
	fmt.Println("d=", d)

}
