/*
	函数类型:
		在Go语言中, 函数也是一种数据类型, 可以通过type来定义它, 函数的类型就是所有拥有相同参数, 相同返回值的一种类型
		函数类型声明:type FuncType func(参数1 数据类型, 参数2 数据类型, ...) (返回值1 返回值类型, 返回值2 返回值类型, ...)
		声明过的函数类型, 可以作为函数形参传入, 在函数内部调用
		例:
			type FuncType func(num int)(int)
			// 定义求和函数
			func sum(num, sum int) (int) {
				for i := 0; i <=num; i++ {
					sum += i
				}
				return sum
			}
			// 将求和函数作为参数传递其他函数
			func FunctypeTest(sum FuncType) {
				var num int
				fmt.Println("请输入要计算的求和数字:")
				fmt.Scan(&num)
				// 调用求和函数
				fmt.Printf("sum=%v", sum(num))
			}
*/

package main

import (
	"fmt"
)

type FuncType func(int) int

type FuncType2 func(int, int) int

// 定义求和函数
func sum(num int) int {
	var sum int
	for i := 0; i <= num; i++ {
		sum += i
	}
	return sum
}

func add(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

// 将求和函数作为参数传递其他函数
func FunctypeTest1(sum FuncType) {
	var num int
	fmt.Println("请输入要计算的求和数字:")
	fmt.Scan(&num)
	// 调用求和函数
	fmt.Printf("sum=%d\n", sum(num))
}

func FunctypeTest2(add, minus FuncType2) {
	var a, b int
	//fmt.Println("请输入变量a和变量b的值:")
	fmt.Print("请输入变量a的值:")
	fmt.Scan(&a)
	fmt.Print("请输入变量b的值:")
	fmt.Scan(&b)
	fmt.Printf("a+b=%d\n", add(a, b))
	fmt.Printf("a-b=%d\n", minus(a, b))
}

func main() {
	FunctypeTest1(sum)
	FunctypeTest2(add, minus)
}
