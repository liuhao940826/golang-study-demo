package main

import (
	"fmt"
)

func main() {

	defer func() {
		fmt.Println("我在这里")
	}()

	go func() {
		pos := "123456"
		defer fmt.Print("我最后打")
		for i := range pos {
			fmt.Println(i)
		}
	}()

}
