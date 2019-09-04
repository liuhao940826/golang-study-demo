package main

import "fmt"

func main() {

	//有多少个方括号就是多少维
	var a [3][5]int

	k := 0
	for i := 0; i < 3; i++ {

		for j := 0; j < 4; j++ {

			k++
			a[i][j] = k
			fmt.Printf("a[%d][%d] = %d ", i, j, a[i][j])

		}
		fmt.Printf("\n")

	}

	//全部初始化
	b := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println("b=", b)

	//部分初始化
	c := [3][4]int{{1, 2, 3, 4}, {5}, {9, 10, 11, 12}}
	fmt.Println("c=", c)

	//指定初始化
	d := [3][4]int{1: {1, 5, 5}, 2: {3, 5, 7}, 0: {6}}
	fmt.Println("d=", d)

}
