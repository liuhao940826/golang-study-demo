package main

//使用指针变量获取命令行的输入信息
import (
	"flag"
	"fmt"
)

// mode 定义一个变量mode 类型根据.String来 是string类型
// 命令行输入的时候变量名叫mode  默认值default value  解释说明 instructions
//flag 返回的是指针变量
var mode = flag.String("mode", "default value", "instructions")

func main() {
	//解析命令行参数
	flag.Parse()
	//输出命令行参数  *mode
	fmt.Println("指针变量的值", *mode)
	//go  run pointer_scan.go  --mode=fast

	//创建指针变量的另一种方法 str 指针变量
	str := new(string)
	//把hello 的赋值给指针变量str的值
	*str = "hello"
	fmt.Println(*str)

}
