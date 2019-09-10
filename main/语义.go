package main

import "fmt"

type Person7 struct {
	name string
	sex  byte
	age  int
}

//接受者为普通变量 值传递
func (tmp Person7) SetValueInfo(name string, sex byte, age int) {
	tmp.name = name
	tmp.age = age
	tmp.sex = sex
	fmt.Printf("值语义 %p  这是不一样的 这个是copy值的地址 \n", &tmp)
}

//接受者为指针变量 引用传递
func (tmp *Person7) SetPointerInfo(name string, sex byte, age int) {

	tmp.name = name
	tmp.age = age
	tmp.sex = sex
	fmt.Printf("指针语义%p \n", tmp)
}

func main() {
	p := Person7{"mike", 'm', 25}
	fmt.Printf("p 的地址 %p p的值 %v \n", &p, p)

	p.SetValueInfo("liuhao", 'f', 222)
	fmt.Println("p ", p)

	(&p).SetPointerInfo("liuhao", 'f', 222)
	fmt.Println("p ", p)
}
