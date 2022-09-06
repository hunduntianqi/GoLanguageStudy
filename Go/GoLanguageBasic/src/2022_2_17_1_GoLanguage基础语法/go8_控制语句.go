/*
	条件语句-if, if...else,嵌套if, if...else
		格式:
			if 判断条件表达式 {
				条件成立要执行的代码
			} else {
				条件不成立执行的代码
			}
	选择语句:switch,select
		1. switch语句用于基于不同条件执行不同的动作, 每一个case分支都是唯一的, 从上至下逐一测试, 直到匹配为止
		2. switch语句执行过程从上到下
		3. switch语句还可以被用于type-switch来判断某个interface变量中存储的变量类型
		格式:
			switch 条件语句 {
				case 匹配值1:
					成功匹配后执行的代码
				case 匹配值2:
					成功匹配后执行的代码
				...
				default:
					所有匹配值都不满足条件后执行的代码
			}
			注意:Go语言中meigecase最后都默认带有break,要想继续执行下面的代码, 可以在case末尾使用fallthrough关键字
	循环语句:for
		格式:
			1. 无条件循环(无线循环)
				for {
					要循环执行的代码
				}
			2. 有条件循环(满足条件循环继续, 不满足条件循环中止)
				for 变量初始值; 循环条件; 变量值控制语句 {
					条件满足执行的代码
				}
				例:
				for i := 1; i < 10; i++ {
					fmt.Println(i)
				}
	控制语句使用的关键字:goto, break,continue
		goto:跳过其他代码块, 直接执行指定代码块,
			goto 指定代码块 在指定代码块之前, 代码块执行一次
			goto 指定代码块 在指定代码块之后, 代码循环执行多次(类似与无限循环)
		break:直接中止循环执行
		continue:跳过本轮循环, 进入下一次循环
*/

package main

import "fmt"

func main() {

	//a := 3
	// if 条件语句使用
	//if a > 3 {
	//	fmt.Println("a的值是大于三的")
	//} else {
	//	fmt.Println("a的值不大于三")
	//}

	// switch选择语句使用
	//switch 5 / 2 {
	//case 1:
	//	fmt.Println("计算值为1")
	//case 2:
	//	fmt.Println("计算值为2")
	//case 3:
	//	fmt.Println("计算值为3")
	//case 4:
	//	fmt.Println("计算值为4")
	//case 5:
	//	fmt.Println("计算值为5")
	//default:
	//	fmt.Println("所有计算结果都不对！！")
	//}
	// switch语句判断变量类型
	//var a interface{} = string("3.7")
	//switch a.(type) {
	//case int:
	//	fmt.Println("变量的类型为整型")
	//case float32:
	//	fmt.Println("变量的类型为浮点型")
	//case string:
	//	fmt.Println("变量的类型为字符串类型")
	//default:
	//	fmt.Println("以上类型都不对！！")
	//}

	//// for 语句使用
	////var number int = 1
	//var start = time.Now()
	//fmt.Println(start)
	//for i := 1; i < 10; i++ {
	//	for j := 1; j <= i; j++ {
	//		fmt.Print(j, "*", i, "=", i*j, "\t")
	//	}
	//	fmt.Println()
	//}
	//var end = time.Now()
	//fmt.Println(end)
	//time.Sleep(10 * time.Second)

	// goto用法
	var i int = 1
One:
	fmt.Println("这是代码块1被执行, i的值为:", i)
	i++
	goto One
}
