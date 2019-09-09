package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

//结构体字段可以通过结构体指针来访问。
//
//通过指针间接的访问是透明的。
func main() {
	v := Vertex{1, 2}
	p := &v
	fmt.Println(v)
	change(p)
	fmt.Println(v)
}

func change(p *Vertex) {
	fmt.Println("指针指向变量值", p.X)
	p.X = 1e9 //指针变量是可以直接获取到它的值
}
