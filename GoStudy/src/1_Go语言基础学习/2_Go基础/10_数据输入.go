package main

import "fmt"

/*
   输入函数:
        fmt.Scanf("格式化字符串", &接收键盘输入信息的变量)
        fmt.Scan(&接收键盘输入信息的变量)
*/

func main() {
	// 定义变量接收键盘输入数据
	fmt.Println("请输入你要说的话:")
	var text string
	fmt.Scan(&text)
	fmt.Println(text)
}
