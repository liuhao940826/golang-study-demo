package main

import "fmt"

func main() {

	total := int64(700)

	size := 100

	totalPage := 0

	if int(total)%size == 0 {
		totalPage = int(total) / size
	} else {
		totalPage = int(total)/size + 1
	}

	fmt.Println("总页数:", totalPage)

}
