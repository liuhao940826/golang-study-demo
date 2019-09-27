package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	resp, err := http.Get("http://example.com/")

	if err != nil {
		// handle error
		defer resp.Body.Close()

		body, err2 := ioutil.ReadAll(resp.Body)

		if err2 != nil {
			fmt.Println("读取响应体失败......")
		}

		fmt.Printf("body %+v", body)
	} else {
		fmt.Println("请求失败.....")
	}

}
