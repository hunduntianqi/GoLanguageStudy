/*
	select:
		Go语言中提供了关键字select, 通过select可以监听channel上的数据流动
		select用法与switch类似, 由select开始一个新的选择块, 每个选择条件由case语句描述
		与switch语句相比, select语句有较多的限制, 其中最大的一条限制为: 每个case语句里必须是一个IO操作
			select语句结构:
				select {
					case <- chan1:
						// chan1成功读取数据后执行的内容
					case chan2 <- 1:
						// 成功向chan2写入数据后执行的内容
					default:
						// 以上条件均不满足后执行的内容
				}
		在一个select语句中, Go语言会按顺序从头至尾评估每一个发送和接收的语句, 如果其中的任意一条语句可以继续执行(没有被阻塞),
		就从这些可以执行的语句中任意选择一条来使用
		如果没有任意一条语句可以执行(所有case语句均被阻塞), 存在两种可能:
			1. 有default语句: 会执行default语句内容, 同时程序会从select语句后的语句中恢复
			2. 没有default语句: select语句将被阻塞, 直到至少有一个case语句可以进行下去
*/

package main

import "fmt"

// 定义生产者函数, ch只写, quit只读
func fibonaqi(ch chan<- int, quit <-chan bool) {
	x, y := 1, 1
	for {
		// 监听channel数据流动
		select {
		case ch <- x:
			x, y = y, x+y
		case flag := <-quit:
			fmt.Println("flag =", flag)
			return
		}
	}
}
func main() {
	/*
		使用select实现斐波那契数列:
			即从第三个数字开始, 每一个数字都等于前两个数字相加
			1 1 2 3 5 8 13 21 ...
	*/
	// 定义channel保存数据
	ch := make(chan int)    // 进行数据通信
	quit := make(chan bool) // 判断程序是否结束
	// 消费者, 从channel中读取数据
	// 新建协程
	go func() {
		for i := 0; i < 1000; i++ {
			num := <-ch
			fmt.Println(num)
		}
		// 可以停止
		quit <- true
	}()
	// 生产者, 向channel中写入数据
	fibonaqi(ch, quit)
}
