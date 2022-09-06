/*
	defer关键字作用:
		用于延迟调用一个函数或者方法(或者当前所创建的匿名函数)
		注意:defer语句只能出现在函数或方法的内部
		使用格式:
		函数名 func() {
			代码1执行
			代码2执行
			defer 代码3执行 //此语句会在函数执行结束前调用, 函数会先执行此语句后面的代码
		}
		多个的defer语句执行顺序:
			如果一个函数中有多个defer语句, 会以LIFO(后进先出, 从下往上)的顺序执行, 即使函数或某个延迟调用发生错误, 这些调用
		仍然会被执行
	defer语句和匿名函数结合执行:
		defer func() {
			功能代码
		} ()
		匿名函数的执行在其他defer修饰代码之前
*/

package main

import "fmt"

func main() {
	defer fmt.Println("我叫郭鹏涛")
	fmt.Println("我爱陈欣妮😗")
	defer fmt.Println("我的老婆陈欣妮最可爱😗")

	defer func() {
		fmt.Println("这是一个匿名函数")
	}()
}
