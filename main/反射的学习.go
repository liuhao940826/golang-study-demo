package main

import (
	"fmt"
	"reflect"
)

type reflectType struct {
	name string
	id   int
}

func main() {

	r := reflectType{}

	t := reflect.TypeOf(r)

	fmt.Println("t,name=", t.Name())
	fmt.Println("t,kind=", t.Kind())

}
