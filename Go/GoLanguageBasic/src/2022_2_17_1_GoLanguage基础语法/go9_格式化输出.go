/*
	格式化输出函数
		fmt.Printf(格式化字符串, 变量)
*/

package main

import "fmt"

func main() {
	a := 10
	b := "abc"
	c := 'a'
	d := 3.14
	// %T操作变量所属类型
	fmt.Printf("%T, %T, %T, %T\n", a, b, c, d)
	// 精确格式化
	fmt.Printf("%d, %s, %c, %F\n", a, b, c, d)
	// 万能格式化%v
	fmt.Printf("%v, %v, %v, %v", a, b, c, d)
}
