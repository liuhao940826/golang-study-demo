package main

import (
	"fmt"
)

func main() {
	//
	//
	//resp, err := http.Get("http://example.com/")
	//...
	//resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
	//...
	//resp, err := http.PostForm("http://example.com/form",
	//	url.Values{"key": {"Value"}, "id": {"123"}})

	var orgId = 17

	str := fmt.Sprintf("polaris:org_%d:priority_list", orgId)

	fmt.Println("字符串:", str)
}
