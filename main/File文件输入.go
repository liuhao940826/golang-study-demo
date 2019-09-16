package main

import (
	"fmt"
	"os"
)

func main() {

	path := "./demo.txt"

	writeFile(path)

}

func writeFile(path string) {
	f, err := os.Create(path)

	if err != nil {
		fmt.Println("err=", err)
		return
	}

	//使用完毕关闭
	defer f.Close()

	var buf string

	for i := 0; i < 10; i++ {

		buf = fmt.Sprintf("i=%d\n", i)

		n, err := f.WriteString(buf)

		if err != nil {
			fmt.Println(buf, "err", err)
		}

		fmt.Println(buf, "n", n)
	}

}
