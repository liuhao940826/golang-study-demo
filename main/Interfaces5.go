package main

import "fmt"

//空接口就是一个万能类型可以保存任意接口的值
type AnyThing interface {
}

type Bus struct {
}

type Car struct {
}

func main() {

	var i interface{} = 1

	fmt.Println("i=", i)

	i = "abc"

	fmt.Println("i=", i)

}
