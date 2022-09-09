package main

import "fmt"

/*
   函数:
        是Go语言的基本代码块
        函数定义:
            func 函数名 (形参) (返回值类型) {
                函数体
            }
            func: 函数声明由func关键字定义
        函数调用:
            无返回值:
                函数名 (传入参数)
            有返回值:
                变量 = 函数名(传入参数)
                注意: go语言中支持多返回值, 所以用来接收返回值的变量可能不止一个
		函数参数传递方式:
			1. 值传递: 指在调用函数时, 只将实参的值传入函数, 在函数内部对形参的操作不会影响外部实参
			2. 引用传递: 在调用函数时, 将实参的地址值传递进函数中, 在函数内部的操作将会同步改变外部实参的值
			默认情况下, go语言使用的是值传递
*/

func main() {
	// 定义两个int变量, 调用函数输出最大值
	var num1 = 20
	var num2 = 30
	fmt.Println("num1 和 num2之间的最大值是:", max(num1, num2))
	// 调用多返回值函数
	name, age := swap("郭鹏涛", 24)
	fmt.Println("name:", name, "age:", age)
	nextNumber := getSequence()
	for i := 0; i < 3; i++ {
		fmt.Println(nextNumber())
	}
}

// 定义函数判断最值
func max(num1, num2 int) int {
	// 定义中间变量
	var temp int
	if num1 > num2 {
		temp = num1
	} else {
		temp = num2
	}
	return temp
}

// 定义多返回值函数
func swap(name string, age int) (string, int) {
	return name, age
}

// 定义闭包函数
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
