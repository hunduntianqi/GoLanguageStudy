/*
	匿名组合:
		体现面向对象的继承特性
		匿名字段:
			一般情况下, 定义结构体时字段名与其类型一一对应, 实际上Go支持只提供类型,而不写字段名的方式,这就是匿名字段,也成为嵌入字段
		成员操作:
			与结构体一致
		同名字段:
			指结构体与其内部嵌套结构体匿名字段中由相同的字段名
			同名字段操作原则:就近原则, 现在自己结构体中找到对应字段进行操作, 找不到再去嵌套字段中找对应字段进行操作
		其他匿名字段:
			非结构体匿名字段:
				type Student struct {
					Person // 结构体匿名字段
					int // 基础类型匿名字段
				}
			结构体指针类型匿名字段:
				type Student struct {
					*Person // 结构体指针匿名字段
					int // 基础类型匿名字段
				}
				指针字段初始化
				结构体变量.Person = &Person{匿名字段成员赋值}
*/

package main

import "fmt"

// 定义人结构体
type Person struct {
	name string //名字
	sex  byte   // 性别
	age  int    // 年龄
}

// 定义学生结构体
type Student struct {
	Person // 匿名字段Person, 默认Student结构体拥有Person结构体的全部字段
	id     int
	addr   string
}

func main() {
	var s1 Student
	s1.name = "郭鹏涛"
	s1.sex = 'm'
	s1.age = 24
	fmt.Println("s1=", s1)

}
