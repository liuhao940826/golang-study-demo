package main

import "fmt"

func main() {


		var strs ="name";
		target := ""

		for _, data:= range strs{

			var flag = matchASCII(target,data);

			fmt.Println("猜猜我是什么",flag);
		}

		//fmt.Printf("输出的数字:"+target)


		//flag:= true;
		//
		//fmt.Print(flag)

}

func matchASCII(target string,ch int32)(isContinue bool){
	if ch >= '0' && ch <= '9' {
		target += string(ch)
		isContinue=true;
	}
	return
}
