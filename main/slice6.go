package main

import (
	"fmt"
	"math/rand"
	"time"
)

//切片做函数参数 引用传递
func main() {
	//切片和底层数组的关系

	var n int = 10

	s := make([]int, n)

	initDate(s)

	fmt.Println(s)

	BubbleSort(s)

	fmt.Println(s)
}

func BubbleSort(s []int) {

	n := len(s)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {

			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}

}

func initDate(s []int) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(s); i++ {
		s[i] = rand.Intn(100) //100以内随机数
	}

}
