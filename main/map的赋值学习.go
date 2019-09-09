package main

import "fmt"

type Upd1 map[string]interface{}

type ValueU struct {
	id int64
}

func main() {

	upd := Upd1{}
	fmt.Println("upd", upd)
	changeUpd1(&upd)
	fmt.Println("upd", upd)

	upd2 := Upd1{}
	fmt.Println("upd", upd2)
	upd2["test"] = ValueU{id: 1}
	changeUpd2(&upd)
	fmt.Println("upd", upd)

}

func changeUpd1(upd1s *Upd1) {
	(*upd1s)["titile"] = 1
	(*upd1s)["lalala"] = "bababa"
}

func changeUpd2(upd1s *Upd1) {
	(*upd1s)["test"] = 1
}
