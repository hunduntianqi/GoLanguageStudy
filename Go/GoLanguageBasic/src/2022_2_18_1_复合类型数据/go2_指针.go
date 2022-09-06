/*
	指针:
		指一个代表某个内存地址的值,一般为内存中存储的一个变量值的起始位置
	Go语言中指针特点:
		1.默认值为nil, 没有NULL常量
		2.操作符"&"取变量地址, "*" 通过指针访问目标对象
		3.不支持指针运算, 不支持"->"运算符, 直接用"."访问目标成员
	指针的定义:
		var 变量名 *数据类型
*/

package main

import "fmt"

func main() {
	// 每个变量定义后有两个含义:1.变量的内存, 2. 变量的地址
	var a int
	a = 10
	fmt.Printf("a=%d \n", a)  // 取变量内存中的内容
	fmt.Printf("a=%v \n", &a) // 取变量地址
	fmt.Printf("a=%p \n", &a) // 取变量地址
	// 定义指针保存变量地址
	var p *int
	p = &a
	fmt.Println("p=", p)
	// 通过指针修改变量a的值
	*p = 100 // 相当于修改了a的值为100, p为变量a的地址, *p相当于访问变量a
	fmt.Printf("*p=%d\n", *p)
	fmt.Printf("a=%d", a)

}
