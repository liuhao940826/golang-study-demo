package main

import (
	"fmt"
	"runtime"
)

func main() {

	//单核的时候 0000000000 11111111111  双核的时候可能 0000111000011111111000
	gomaxprocs := runtime.GOMAXPROCS(2) //指定单核运算

	fmt.Println("gomaxprocs", gomaxprocs)

	for {

		go fmt.Print(1)
		fmt.Print(0)

	}

}
