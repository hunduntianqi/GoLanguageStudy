/*
	函数定义:
		格式:
			func 函数名(参数列表) (参数对应返回值类型, 没有可不写) {
				函数体(功能代码)
				return 要返回的数据
			}
		函数定义说明:
			func: 函数定义关键字
			函数名: 函数名字母小写开头为private, 大写开头为public
			参数列表: 函数可以有0个或多个参数, 参数格式为: 变量名 变量类型, 如果有多个参数之间用逗号分割, 不支持默认参数
			返回类型:
				1. 返回类型视实际需求而定, 可以没有
				2. 如果只有一个返回值且不声明返回值类型, 可以省略
				3. 如果有返回值, 那么必须在函数的内部添加return语句来返回数据并中断函数
		无参无返回值函数:
			func MyFunc() {
				功能代码
			}
		有参无返回值函数:
			func MyFunc(参数1 参数类型, 参数2 参数类型) {
				功能代码
			}
		不定参数类型函数
			func MyFunc(参数1 参数类型, 参数名称 ...参数类型) {
				功能代码
				fmt.Println("len(args)=", len(args)) // 可以打印出用户传参个数
			}
			注意: 固定参数一定要传参,不定参数可以不传参, 而且不定参数必须放在形参列表最后一个参数
			len()函数计算中文一个字符为3
		无参有一个返回值
			func MyFunc() (int) {
				return 返回数据
			}
		无参多个返回值函数
			func MyFunc() (返回值1 数据类型, 返回值2 数据类型, ...) {
				return 返回值1, 返回值2, ...
			}
		带参数有返回值函数
			func MyFunc(参数1 数据类型, 参数2 数据类型,...) (返回值1 数据类型, 返回值2 数据类型, ...) {
				return 返回值1, 返回值2,...
			}

*/

package main

import "fmt"

// 定义无参无返回值函数
func MyFunc1() {
	fmt.Println("这是一个无参无返回值函数")
}

// 定义带参数函数
func MyFunc2(num int, name string, fl float32) {
	fmt.Printf("%d, %s, %0.2f\n", num, name, fl)
}

// 定义不定长参数函数
func MyFunc3(name string, args ...int) {
	fmt.Printf("用户传参个数为len(args+name)=%d\n", len(args)+len(name))
}

// 定义无参一个返回值
func MyFunc4() int {
	var num int
	num = 3
	return num
}

// 定义无参多个返回值
func MyFunc5() (int, string, float32) {

	return 5, "郭鹏涛", 3.14
}

// 定义带参数有返回值函数
func MyFunc6(num int, str string) (int, string) {
	fmt.Printf("请输入变量num的值:")
	fmt.Scan(&num)
	fmt.Printf("请输入变量str的值:")
	fmt.Scan(&str)

	return num, str
}

func main() {
	MyFunc1()
	MyFunc2(10, "郭鹏涛", 3.14)
	MyFunc3("郭鹏涛", 1, 2, 3, 4)
	// 无参带返回值函数调用
	num := MyFunc4()
	fmt.Println("num =", num)
	// 多返回值函数调用, 需要多个变量接收返回数据
	a, b, c := MyFunc5()
	fmt.Printf("a=%d, b=%s, c=%0.2f\n", a, b, c)
	// 有返回值有参数函数调用
	Inum, Sname := MyFunc6(0, "")
	fmt.Printf("变量Inum=%d, 变量Iname=%s\n", Inum, Sname)
}
