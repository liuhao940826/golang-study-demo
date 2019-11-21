package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	str1 := []string{"one", "two", "three"}
	str2 := []string{"one", "two", "three"}
	str3 := []string{"one", "two", "three"}

	result := &[]interface{}{str2, str1, str3}

	fmt.Println("result", result)

	bytes, _ := json.Marshal(result)

	fmt.Println("json数组", string(bytes))

	//json转数组

	sstruct := &[]interface{}{}

	_ = json.Unmarshal(bytes, sstruct)

	fmt.Println("sstruct", *sstruct)

}

//func main() {
//	s := []interface{}{1, 2, 3, 4, 5, 6, 7, 8}
//HERE:
//	for k, v := range s {
//		if v == 4 || v == 6 || v == 7 {
//			// temp := s[k+1:]
//			fmt.Println("k", k)
//
//			temp := make([]interface{}, 0)
//			for _, v := range s[k+1:] {
//				temp = append(temp, v)
//			}
//
//			fmt.Println("temp1", temp)
//
//			s = append(s[:k], "(")
//			fmt.Println(s)
//			s = append(s, []interface{}{fmt.Sprintf("%d", v), fmt.Sprintf("%d", v)}...)
//			fmt.Println(s)
//			s = append(s, ")")
//			fmt.Println(s)
//			fmt.Println("temp2", temp)
//			s = append(s, temp...)
//			fmt.Println(s)
//			goto HERE
//		}
//	}
//}
