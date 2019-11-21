package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	v1 := TestAppend{1, 1}
	v2 := TestAppend{1, 2}
	v3 := TestAppend{1, 3}

	list := []*TestAppend{}

	list = append(list, &v1)
	list = append(list, &v2)
	list = append(list, &v3)

	bytes1, _ := json.Marshal(list)
	fmt.Println("自增之前的list", string(bytes1))

	for i := 0; i < len(list); i++ {
		if list[i].Id == 3 {
			list[i].Num++
		}

		if list[i].Num >= 2 {
			list = append(list[:i], list[i+1:]...)
			i--
		}
	}

	bytes, _ := json.Marshal(list)

	fmt.Println("自增之后的list", string(bytes))
}

type TestAppend struct {
	Num int
	Id  int
}
