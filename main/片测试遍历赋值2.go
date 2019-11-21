package main

import "fmt"

func main() {

	ids := []int{1, 2, 3, 4}

	handlerDealDailyProjectPushMsgRetry(&ids)

	fmt.Println("最后的id集合:", ids)

}

func handlerDealDailyProjectPushMsgRetry(orgIds *[]int) {

	v1 := TestAppend2{1, 1}
	v2 := TestAppend2{1, 2}
	v3 := TestAppend2{1, 3}

	list := []*TestAppend2{}

	list = append(list, &v1)
	list = append(list, &v2)
	list = append(list, &v3)

	for i := 0; i < len(list); i++ {
		*orgIds = append(*orgIds, list[i].Id)
	}
}

type TestAppend2 struct {
	Num int
	Id  int
}
