package main

import "fmt"

//指针（pointer）概念在 Go 语言中被拆分为两个核心概念：
//类型指针，允许对这个指针类型的数据进行修改。传递数据使用指针，而无须拷贝数据。类型指针不能进行偏移和运算。
//切片，由指向起始元素的原始指针、元素数量和容量组成。

type pointerMode struct {
	Name *string
}

//可以声明指针      指向任何类型的值   来表明它的原始性或结构性；
//你可以在指针类型前面加上 * 号（前缀）来获取指针所指向的内容，
//这里的 * 号是一个类型更改器 使用一个指针引用一个值被称为间接引用
//当一个指针被定义后没有分配到任何变量时，它的值为 nil。一个指针变量通常缩写为 ptr。
//每个变量在运行时都拥有一个地址，这个地址代表变量在内存中的位置。Go 语言中使用&作符放在变量前面对变量进行“取地址”操作。

func main() {
	//随便定义一个int变量
	v := 1
	// *****其中 v 代表被取地址的变量，被取地址的 v 使用 ptr 变量进行接收，
	// *****ptr  代表指针变量,     指针变量(ptr)的类型为*T
	// *代表指针 T代表指针类型(重要事情强调三遍)
	// *代表指针 T代表指针类型(重要事情强调三遍)
	// *代表指针 T代表指针类型(重要事情强调三遍)
	// *****&   :  使用&    作符放在变量前面对变量    进行“取地址”操作。(互补)
	// *****      取地址操作符&和取值操作符*是一对互补操作符，&取出地址，*根据地址取出地址指向的值。(互补)
	ptr := &v // v的类型为T
	//这边的ptr 的地址是等于 v的地址的  %p输出指针地址
	fmt.Printf("ptr的类型%T  ptr的这个指针变量中存放的v的地址值%v", ptr, ptr)

	fmt.Println("==================================")
	str := "banana"
	fmt.Printf(" 内容 %s \n 类型%T\n 地址%p \n", str, &str, &str)

	pm := pointerMode{&str}
	pm1 := pointerMode{&str}

	fmt.Printf("实体类型%T\n", &pm)

	fmt.Printf("获取地址pm %p  pm中name的值 %v \n", &pm.Name, *pm.Name)

	fmt.Printf("获取地址pm1 %p pm1中name的值 %v \n", &pm1.Name, *pm1.Name)
	//
	//fmt.Printf("%p %p %p\n", &v, &str, ptr)
	//fmt.Println("%p %p %p", &v == ptr)

	//******输出值在每次运行是不同的，代表 cat 和 str 两个变量在运行时的地址。

	//******在 32 位平台上，将是 32 位地址；64 位平台上是 64 位地址。

	//******提示：变量、指针和地址三者的关系是：每个变量都拥有地址，指针的值就是地址。

	//从指针获取指针指向的值

}
