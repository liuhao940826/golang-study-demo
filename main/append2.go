package main

import (
	"fmt"
	"strconv"
	"strings"
)

type appenStruct struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

//*操作符作为右值时，意义是取指针的值；作为左值时，也就是放在赋值操作符的左边时，表示 a 指向的变量。
// 其实归纳起来，*操作符的根本意义就是操作指针指向的变量。当操作在右值时，就是取指向变量的值；当操作在左值时，就是将值设置给指向的变量。

//传进来地址
//func changeInfo(list *[]*appenStruct){
//	//* list 代表获取 指针变量list 对应的地址的值去遍历
//	for _,data := range *list{
//		fmt.Printf("获取的类型 %T %v \n",*data,*data)
//		//这边的* 标识获取指针变量
//		(*data).Id=333
//		fmt.Println("change data ",*data)
//	}
//}
//
//
////数组长度在定义后无法再次修改 数组是值类型 每次传递产生一个副本  用slice来弥补不足
//func main() {
//
//	//append 在切片末尾追加数字 如果不够就扩容  如果超过原来的容量通常以3倍扩容
//	s := make([]*appenStruct, 0, 1)//[]
//
//	fmt.Println("s",s)
//
//	s = append(s, &appenStruct{Id:1})//[{id:1}]
//
//	fmt.Println("s",s)
//	for _,data := range s{
//		//这边的* 标识获取指针变量
//		fmt.Println("out data ",*data)
//	}
//
//
//	changeInfo(&s)
//	fmt.Println("s",s)
//
//	for _,data := range s{
//		//这边的* 标识获取指针变量
//		fmt.Println("out data ",*data)
//	}
//
//}

//func changeInfo(list *[]appenStruct){
//	//* list 代表获取 指针变量list 对应的地址的值去遍历
//	for _,data := range *list{
//		fmt.Printf("获取的类型 %T %v \n",data,data)
//		//这边的* 标识获取指针变量
//		data.Id=333
//		fmt.Println("change data ",data)
//	}
//}
//
//
////数组长度在定义后无法再次修改 数组是值类型 每次传递产生一个副本  用slice来弥补不足
//func main() {
//
//	//append 在切片末尾追加数字 如果不够就扩容  如果超过原来的容量通常以3倍扩容
//	s := make([]appenStruct, 0, 1)//[]
//
//	fmt.Println("s",s)
//
//	s = append(s, appenStruct{Id:1})//[{id:1}]
//
//	fmt.Println("s",s)
//
//	changeInfo(&s)
//	fmt.Println("s",s)
//
//}

func appendStr(str *string, ints []int64) {

	for _, v := range ints {
		*str += strconv.FormatInt(v, 10)
		*str += ","
	}
	*str = strings.TrimRight(*str, ",")

}

func main() {

	//append 在切片末尾追加数字 如果不够就扩容  如果超过原来的容量通常以3倍扩容
	slice := []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	str := ""
	fmt.Println("ssss=", str)
	appendStr(&str, slice)

	fmt.Println("s=", str)
}
