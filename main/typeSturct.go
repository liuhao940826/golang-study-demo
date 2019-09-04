package main

import "fmt"

type person struct {
	//姓名
	Name string
	//年龄
	Age int
}

type dog struct {
	//姓名
	string
	//年龄
	Age int
}

func main() {

	person1 := person{"liuhao", 25}
	fmt.Println(person1.Name)

	dog1 := dog{string: "旺财"}

	fmt.Println(dog1.string)

}
