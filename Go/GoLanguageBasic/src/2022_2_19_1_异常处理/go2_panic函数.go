/*
	panic函数:
		一般而言, 当panic异常发生时, 程序会中断运行, 并立即执行再该goroutine(可以理解为线程)中被延迟的函数(defer机制), 然后程序崩溃
	并输出日志信息, 日志信息包括panic value和函数调用的堆栈跟踪信息
	注意: 不是所有的panic异常都来自运行时, 直接调用内置的panic函数也会引发panic异常
	panic函数接受任何值作为参数:
		func panic(v interface{}) {}



*/

package main

func main() {

}
