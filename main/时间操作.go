package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	//Date返回一个时区为loc、当地时间为：
	currentTime := time.Now()
	endTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 0, 0, currentTime.Location())
	fmt.Println(endTime)
	fmt.Println(endTime.Format("2006-01-02 15:04:05"))

	hour := -1
	itoa := strconv.Itoa(hour)

	fmt.Println("转换后的结构", itoa)

}
