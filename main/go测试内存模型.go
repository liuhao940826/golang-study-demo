package main

import "fmt"

func main() {

	a := 1

	go func() {
		a += 1
		fmt.Println("a被执行了...............")
	}()

	go func() {
		a += 1
		fmt.Println("b被执行了...............")
	}()

	fmt.Println("a的值:", a)

}
