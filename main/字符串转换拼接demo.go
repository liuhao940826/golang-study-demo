package main

import (
	"fmt"
	"strconv"
)

func main() {

	secondInterval := 600

	timeStr := fmt.Sprintf("%s"+strconv.Itoa(secondInterval)+"%s", "-", IntervalSecondUnit)

	fmt.Println("转化你的结果:", timeStr)

}
