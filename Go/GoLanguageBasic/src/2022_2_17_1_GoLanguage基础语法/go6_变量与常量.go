/*
	变量声明初始化和赋值:
		变量的声明格式:var 变量名称 变量类型
		变量赋值: 变量名称 = 值 / 表达式 / 函数等
		声明和赋值同时进行: var 变量名称 变量类型 = 值 / 表达式 / 函数等
		分组声明格式:
			var (
				i int
				j float32
				name string
			)
		声明多个变量和赋值: var a, b, c int = 1, 2, 3或者 a, b := 1, 2
		注意:
			全局变量声明必须使用var关键字,局部变量可以省略
		特殊变量: "_"
		变量的类型转换:
			Go语言中不存在隐士转换, 类型转换必须是显示的, 并且类型转换只能发生在两种兼容类型之间
			类型转换格式: 变量名称 := 目标类型(需要转换的变量)
		变量可见性规则:
			1. 公有变量: 大写字母开头的变量是可导出的, 指其他包可以读取的变量
			2. 私有变量: 小写字母开头的, 不可导出的变量
		查看变量类型:reflect.TypeOf(变量名称)
		自动推导变量类型 变量名 := 变量值(只能在函数体内部使用)
		局部变量:
			在{}内部定义的变量, 只在{}内部有效
		全局变量:
			定义在函数外部的变量,全局变量必须使用var关键字来声明, 不能使用自动推导类型来定义
	常量定义:
		常量定义形式:
			显示定义:const 常量名称 常量类型 = 常量值
			隐式定义: const 常量名称 = 常量值(通常叫无类型常量)
			组合定义:
			const (
				常量名 常量类型 = 常量值
				常量名 = 常量值
			)
			定义多个常量:
			const 常量1, 常量2, ...常量类型 = 常量值1, 常量值2, ...
			注意:
				1.常量可以使用内置表达式定义, 例如: len(), unsafe.Sizeof()等
		常量支持类型:常量目前只支持布尔型, 数字型, 和字符串型
		特殊常量-iota的使用:
			iota在const关键字出现时将被重置为0
			const中每新增一行常量声明,iota将计数一次
			iota常见使用方法:
				1.跳值使用法
				2.插队使用法
				3.表达式隐式使用法
				4.单行使用法
*/

package main

import "fmt"

//// 声明单个变量
//var num int = 10
//
//// 变量赋值
//var b string = "郭鹏涛"
//
//// 多个变量同时赋值
//var c, d int = 10, 20
//
//// 变量的分组赋值
//var (
//	i    int     = 10
//	f    float32 = 2.75
//	name string  = "老婆陈欣妮"
//)

const acon = iota
const (
	bcon = iota
	ccon = iota
)

func main() {
	//// 变量的分组赋值
	//var (
	//	i    int     = 10
	//	f    float32 = 2.75
	//	name string  = "老婆陈欣妮"
	//)
	//// 变量类型转换
	//fl := int(f)
	//str := "郭鹏涛最爱陈欣妮"
	//fmt.Println(num)
	//fmt.Println(b)
	//fmt.Println("变量c的值为:", c, "变量d的值为:", d)
	//fmt.Println(i, f, name)
	//fmt.Println("name变量的类型为", reflect.TypeOf(name))
	//fmt.Println(str, "\n", fl)
	fmt.Println("a的常量值为:", acon)
	fmt.Println("b的常量值为:", bcon)
	fmt.Println("c的常量值为:", ccon)
}
