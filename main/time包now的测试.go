package main

import (
	"fmt"
	"time"
)

func main() {

	var a int = 1

	var b uint64 = 1

	fmt.Println("你们相等么", a == int(b))

	fmt.Println("现在时间", time.Now())
}
