/*
	goroutine:
		是Go语言并行设计的核心, 本质上指协程, 是比线程更小的单位, goroutine比thread更易用, 更高效, 更轻便
	goroutine创建:
		在Go语言中, 只需要在函数调用语句前加go关键字, 就可以创建并发执行单元
		当一个程序启动时, 主函数main()在一个单独的goroutine中运行, 被称作main goroutine,
		新的goroutine会用 go语句来创建
		注意:
			如果主协程退出, 其他子协程也会跟着退出执行
			解决方法:
				sync.WaitGroup类型:
				使用方法:
					1. 定义WaitGroup变量
						var wg WaitGroup
					2. 调用Add()函数添加计数
					3. Done()函数减去一个计数
					3. 调用Wait()函数等待子协程执行完毕
*/
package main

import "fmt"

func Task() {
	for i := 0; i < 10; i++ {
		fmt.Println("子协程i=", i)
	}
}

func main() {
	go Task()
	//for i := 0; i < 10; i++ {
	//	fmt.Println("主协程i=", i)
	//}
}
