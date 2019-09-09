package main

import "fmt"

type Test struct {
	A int64
	B string
}

func main(){
	// 引用类型
	map_param := make(map[int64]string,0)
	fmt.Println(map_param)
	mapfunc(map_param)
	fmt.Println(map_param)

	slice_param := make([]int64,0)
	slice_param = append(slice_param,999)
	fmt.Println(slice_param)
	slicefunc(slice_param)
	fmt.Println(slice_param)


	// 值类型
	arr_param := [5]int64{1,2,3,4}
	fmt.Println(arr_param) // 1,2,3,4
	arrfunc(arr_param)
	fmt.Println(arr_param) // 1,2,3,4
	test := Test{A:1,B:"ooo",}
	fmt.Println(test) // 1,ooo
	testfunc(test)
	fmt.Println(test) // 1,ooo
}

func mapfunc(strings map[int64]string) {

	strings[]

}

func slicefunc(param []int64){
	param[0]=10000
}

func testfunc(a Test){
	a.A=9999
	a.B="hah"
}

func arrfunc(param [5]int64){
	param[0]=10000
}