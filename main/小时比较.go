package main

import (
	"fmt"
	"strconv"
	"time"
)

const AppDateFormat = "2006-01-02"
const AppTimeFormat = "2006-01-02 15:04:05"
const AppTimeNotSecondeFormat = "2006-01-02 15:04"
const AppSystemTimeFormat = "2006-01-02T15:04:05Z"
const AppSystemTimeFormat8 = "2006-01-02T15:04:05+08:00"

const (
	IntervalSecondUnit = "s"
	IntervalMinuteUnit = "m"
	IntervalHourUnit   = "h"
)

const (
	Plus  = "+"
	Minus = "-"
)

func main() {
	//获取当天的23:59:00
	currentTime := time.Now()

	//fmt.Println("当前时间", currentTime)
	//
	//dailyEndTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 0, 0, currentTime.Location())
	//
	////获取扫描区间
	//secondInterval := 1800
	//
	////获取每日23:59减少区间后的时间
	//subConverTime := AssemblyDateTime(secondInterval, dailyEndTime, Minus, IntervalSecondUnit)
	//
	//fmt.Println("每日倒计时减少后的时间", subConverTime)
	//
	//convertTime := AssemblyDateTime(secondInterval, currentTime, Plus, IntervalSecondUnit)
	//
	//fmt.Println("增加后的时间", convertTime)

	dateTime := AssemblyDateTime(0, currentTime, Minus, IntervalSecondUnit)

	fmt.Println("dateTime:", dateTime)

	parse, _ := time.Parse(AppTimeNotSecondeFormat, "2019-11-21 00:04")

	convertMinusTime := AssemblyDateTime(10, parse, Minus, IntervalMinuteUnit)
	//获取当前时间小时+分钟
	fmt.Println("不够减的区间时间", convertMinusTime)

	fmt.Println("测试空串大小", "" > "0")

}

func CTime(t time.Time, time_str string) time.Time {
	time_part, err := time.ParseDuration(time_str)
	if err != nil {
		fmt.Println(err)
		return t
	}
	return t.Add(time_part)
}

func AssemblyDateTime(secondInterval int, orgTime time.Time, symbol string, intervalUnit string) string {
	//当前时间加上时间区间的时间
	plusTimeStr := fmt.Sprintf("%s"+strconv.Itoa(secondInterval)+"%s", symbol, intervalUnit)

	fmt.Println("表达式", plusTimeStr)
	currentTimeAddInterval := CTime(orgTime, plusTimeStr)
	//转换时间区间的字符串时间 =  当前时间+ 时间区间
	dateTime := strconv.Itoa(currentTimeAddInterval.Hour()) + ":" + strconv.Itoa(currentTimeAddInterval.Minute())
	return dateTime
}
