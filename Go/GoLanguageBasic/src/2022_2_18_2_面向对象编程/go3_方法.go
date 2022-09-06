/*
	方法:
		体现面向对象的封装特性！！
		本质上一个方法是一个和特殊类型关联的函数
		在Go语言中, 可以给任意自定义类型(包括内置类型, 但不包括指针类型)添加相应的方法
		方法定义格式:
			func (receiver ReceiverType) funcName(parameters) (results) {
				功能代码
			}
			注意:
				1. 参数receiver可以任意命名, 如果方法中未使用, 可以省略不写
				2. 不支持重载方法, 不能定义名字相同但是参数不同的的方法
			例:
				type int_test int
				// 给int_test类型绑定一个函数
				func (num int_test) add(b int_test) int_test {
					return num + b
				}
				func main() {
					var num int_test
					num = 10
					fmt.Println(num.add(100))
				}
		匿名字段:
			方法的继承:
				假设匿名字段实现了一个方法, 继承这个匿名字段的结构体同时继承了该字段的成员和方法
			方法的重写(同名方法):
				例:
					结构体Student继承了匿名字段Person, Student实现了一个方法, Person中也实现了一个方法, 两个方法名字相同
				叫做方法重写或同名方法
				重写方法隐式调用特点-就近原则:
					先调用被结构体实现的方法, 找不到再去调用继承的方法
				重写方法显示调用继承重写方法:
					变量名.匿名字段.方法名()
*/

package main

import "fmt"

type int_test int

// 给int_test类型绑定一个函数
func (num int_test) add(b int_test) int_test {
	return num + b
}

func main() {
	var num int_test
	num = 10
	fmt.Println(num.add(100))
}
