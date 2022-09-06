/*
	输入函数:
		fmt.Scanf("格式化字符串", &接受键盘输入信息的变量)
		fmt.Scan(&接受键盘输入信息的变量)
*/

package main

import "fmt"

func main() {
	var a string
	fmt.Println("请输入变量a的值:")
	// 会阻塞等待用户输入
	fmt.Scan(&a)
	fmt.Println("变量a的值为:", a)
}
