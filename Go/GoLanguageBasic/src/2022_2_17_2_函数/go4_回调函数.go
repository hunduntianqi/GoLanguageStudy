/*
	回调函数:
		指有一个参数为函数类型的函数就是回调函数
*/

package main

import "fmt"

type FuncType_test func(int, int) int

// 定义函数相加
func jia(a, b int) int {
	return a + b
}

// 定义函数相减
func jian(a, b int) int {
	return a - b
}

// 定义函数相乘
func multiply(a, b int) int {
	return a * b
}
func divide(a, b int) int {
	return a / b
}

func Test(a, b int, fTest FuncType_test) int {
	return fTest(a, b)
}

func main() {
	a, b := 5, 3

	fmt.Printf("a+b=%d\n", Test(a, b, jia))
	fmt.Printf("a-b=%d\n", Test(a, b, jian))
	fmt.Printf("a * b = %d\n", Test(a, b, multiply))
	fmt.Printf("a / b = %d\n", Test(a, b, divide))
}
