package main

import "fmt"

const (
	DefalutCountryId  = 0
	DefalutStatesId   = 0
	DefalutCityId     = 0
	DefalutIndustryId = 0
)

func main() {

	testID := int64(0)

	fmt.Println("是否相等", testID == DefalutCountryId)

}
