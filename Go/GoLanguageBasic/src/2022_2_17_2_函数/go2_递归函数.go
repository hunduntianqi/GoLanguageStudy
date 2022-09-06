/*
	递归函数:
		指函数可以直接或间接的调用自身
		递归函数通常具有相同的结构:
			跳出条件:根据传入的参数条件判断是否需要停止递归
			递归体:函数要执行的功能代码
*/

package main

import "fmt"

// 定义简单的递归函数计算指定数据累加
func Mysum(num int) int {
	if num == 1 {
		return 1
	}
	return num + Mysum(num-1)
}

func main() {
	// 递归函数调用
	fmt.Println("请输入要计算的累加数字:")
	var sum int
	fmt.Scan(&sum)
	fmt.Printf("sum=%d \n", Mysum(sum))
}
