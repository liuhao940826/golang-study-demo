package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {

	env := "local"

	//读取yaml文件
	v := viper.New()
	//设置读取的配置文件
	v.SetConfigName("application" + "." + env)
	//添加读取的配置文件路径
	//v.AddConfigPath("./main/config/")

	v.AddConfigPath("./config/")
	//windows环境下为%GOPATH，linux环境下为$GOPATH
	//v.AddConfigPath("$GOPATH/src/")
	//设置配置文件类型
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("err:%s\n", err)
	}

	fmt.Printf(
		`
		ServerPort:%s`,
		v.Get("Server.Port"),
	)

	/*
		result:
		TimeStamp:2018-10-18 10:09:23
		CompanyInfomation.Name:Sunny
		CompanyInfomation.Department:[Finance Design Program Sales]
	*/

}
