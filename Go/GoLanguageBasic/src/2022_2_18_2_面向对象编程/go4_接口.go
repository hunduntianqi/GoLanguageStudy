/*
	接口(interface):
		体现面向对象的多态特征
		是一个自定义类型, 接口类型具体描述了一系列方法的集合
		接口类型不能被实例化
		接口定义:
			type 接口名 interface {
				方法1()
				方法2()
			}
			注意:
				1. 接口命名习惯以er结尾
				2. 接口只有方法声明, 没有实现, 没有数据字段
				3. 接口可以匿名嵌入其他接口, 或嵌入到结构中
		接口实现:
			接口是用来定义行为的类型, 这些被定义的行为不直接由接口实现, 而是由用户定义的类型实现, 一个实现了这些方法的具体类型就是
		这个接口类型的实例
		注意:
			实现了接口方法的类型变量可以给接口类型变量赋值
		接口嵌入:
			如果一个interface1作为interface2的一个嵌入字段, 那么interface2隐式的包含了interface1里面的方法
			例:
				type interface1 interface {
					sayHi()
					run()
				}
				type interface2 interface {
					interface1 // 相当于定义了sayHi()方法和run()方法
					study()
				}
		空接口:
			空接口不包含任何方法, 正因为如此, 所有的类型都实现了空接口, 因此空接口可以存储任意类型的数据
			当函数可以接受任意的对象实例时, 我们会将其声明为interface{}, 最典型的例子是标准库fmt中的PrintXXX系列的函数
*/

package main

import "fmt"

// 定义接口
type Humaner interface {
	// 接口中方法定义, 只有方法声明, 没有方法实现, 由其他自定义类型实现
	sayHi()
}

// 自定义学生结构体
type student struct {
	name string
	age  int
}

type Teach struct {
	student
}

// Student实现sayHi()方法
func (s1 *student) sayHi() {
	fmt.Printf("Student[%s, %d]\n", s1.name, s1.age)
}

// Teach实现sayHi()方法
func (t1 *Teach) sayHi() {
	fmt.Printf("Teach[%s, %d]\n", t1.name, t1.age)
}

// 定义函数, 参数为接口, 实现多态特征
func inter(i Humaner) {
	i.sayHi()
}
func main() {
	// 定义接口类型变量
	var i Humaner
	// 实现接口方法的自定义类型变量可以给接口变量赋值
	s := &student{name: "郭鹏涛", age: 24}
	i = s
	i.sayHi()
	t := &Teach{student{
		name: "陈欣妮",
		age:  24,
	}}
	i = t
	i.sayHi()
	// 相同函数, 不同执行结果
	inter(s)
	inter(t)

}
