/*
	结构体:
		是一种聚合的数据类型,由一系列具有相同类型或不同类型的数据构成的数据集合,每个数据称为结构体的成员
		格式:
			type 结构体名 struct {
				数据1 数据类型
				数据2 数据类型
				数据3 数据类型
				数据4 数据类型
				............
			}
		结构体初始化:
			1. 顺序初始化, 每个成员都必须初始化
			2. 指定成员初始化, 未初始化成员自动赋值为0
			3. 结构体指针变量初始化:
				例:
					var s3 *Student = &Student{id: 3, name: "郭鹏涛和陈欣妮", age: 24, sex: "男&女", addr: "河南省"}
					fmt.Println("*s3=", *s3)
				注意:不要忘了取地址操作: &结构体名
		结构体成员的使用:
			1. 操作普通变量-需要使用“.”运算符
			2. 操作结构体变量
				2.1 使用(*指针变量).成员名称或指针变量.成员名称操作指针成员变量
					例:
						// 结构体指针变量初始化
						var s3 *Student = &Student{id: 3, name: "郭鹏涛和陈欣妮", age: 24, sex: "男&女", addr: "河南省"}
						fmt.Println("*s3=", *s3)
						// 结构体成员使用-指针变量
						fmt.Println("s3-name=", (*s3).name)
						fmt.Println("s3-name=", s3.name)
		结构体的比较:
			同类型的两个结构体变量可以相互比较, 比较结果为布尔值, 相同为true, 不同为false
		结构体赋值:
			同类型的结构体可以互相赋值
		结构体做函数参数-值传递:
			func 函数名(参数名称 结构体类型) {

			}
		结构体做函数参数-引用传递:
			引用传递传入参数为取地址类型传参-&结构体变量名称
			func 函数名(参数名称 *结构体类型) {

			}





*/

package main

import "fmt"

// 定义学生信息结构体
type Student struct {
	id   int
	name string
	age  int
	sex  string
	addr string
}

func main() {
	// 结构体顺序初始化, 每个成员都必须初始化
	var s1 Student = Student{1, "郭鹏涛", 24, "男", "河南省洛阳市"}
	fmt.Println(s1)
	// 指定成员初始化,未初始化成员自动赋值为0
	var s2 Student = Student{id: 2, name: "陈欣妮", age: 24}
	fmt.Println(s2)
	// 结构体指针变量初始化
	var s3 *Student = &Student{id: 3, name: "郭鹏涛和陈欣妮", age: 24, sex: "男&女", addr: "河南省"}
	fmt.Println("*s3=", *s3)
	// 结构体成员的使用-普通变量
	// 操作成员需要使用“.”运算符
	fmt.Println("s1-name=", s1.name)
	fmt.Println("s1-age=", s1.age)
	fmt.Println("s1-sex=", s1.sex)
	fmt.Println("s1-add=", s1.addr)
	// 结构体成员使用-指针变量
	fmt.Println("s3-name=", (*s3).name)
	fmt.Println("s3-name=", s3.name)

	// 结构体的比较
	fmt.Println(s1 == s2)
}
