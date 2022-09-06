/*
	goroutine奉行通过通信来共享内存, 而不是共享内存来通信
	引用类型channel:
		是CSP模式的具体体现, 用于多个goroutine通讯, 其内部实现了同步, 确保并发安全
		channel:和map类似, 也是一个对应make创建的底层数据结构的引用, 和其他引用类型一样, channel
		的零值也是nil, 当复制一个channel或用于参数传递时, 只是拷贝了一个channel引用, 调用者和被调
		用者将引用同一个channel对象
		定义channel:
			ch := make(chan type, capacity)
			chan:指定义channel类型变量
			type:指该channel对象中的数据类型
			capacity:当小于0时, channel是无缓冲阻塞读写的
						即:当管道中无数据时, 取出数据操作会造成阻塞
						   当管道中有数据时, 写入数据操作会造成阻塞
					 当大于0时, channel有缓冲, 是非阻塞的, 直到写满capacity个元素才阻塞写入, 当管道中数据被全部取出后
					 再次取数据会造成阻塞
		channel通过操作符<-接收和发送数据, 发送数据和接收数据语法:
			channel <- value: 将value存入管道channel
			<-channel: 从管道中取出数据并丢弃
			x := <-channel: 从管道channel中取出数据并赋值给x
			x, ok := <- channel: 从管道channel中取出数据赋值给x, 并检查管道是否为空或管道
								 是否已关闭
		关闭channel:close(ch) (当不需要再写数据时, 可以关闭数据), 注意:当channel已关闭时,
					无法再写入数据, 但可以取出数据
		使用range遍历管道内容:
			遍历结束或会自动关闭循环
			for num := range ch{
			}
		单方向channel:
			默认情况下, channel是双向的, 既可以向channel中写入数据, 也可以从channel中取出数据
			单向channel变量声明:
				var ch1 chan<- type: 单向channel, 可以写入type类型的数据
				var ch1 <-chan type: 单向channel, 可以取出type类型的数据
				chan<-: 表示数据存入管道
				<-chan: 表示将数据从管道中取出
			双向channel转换为单向channel:
				c := make(chan type)
				var send chan<- type = c: 只能将数据存入管道
				var send <-chan type = c: 只能从管道中取出数据
				注意: 单向channel不能再转换为普通的双向channel
			单向channel应用:作为函数参数传递
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	//// 创建无缓冲channel
	//ch := make(chan string, 0)
	//defer fmt.Println("主协程结束运行！！")
	//go func() {
	//	defer fmt.Println("子携程被调用")
	//
	//	for i := 0; i < 10; i++ {
	//		fmt.Println("子协程 i =", i)
	//		ch <- strconv.FormatInt(int64(i), 10)
	//		//time.Sleep(1 * time.Second)
	//		fmt.Printf("len(ch)=%d, cap(ch)=%d\n", len(ch), cap(ch))
	//	}
	//	//ch <- "子协程工作完毕" // 子协程函数给channel写入数据
	//}()
	////time.Sleep(2 * time.Second)
	//for j := 0; j < 10; j++ {
	//	num := <-ch
	//	fmt.Println("num =", num)
	//}
	////<-ch
	////<-ch // channel没有数据, 此处会阻塞

	// 创建有缓存channel
	ch_buf := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			ch_buf <- i
			fmt.Printf("子协程 i = %d \n", i)

			fmt.Printf("len(ch_buf)=%d, cap(ch_buf)=%d\n", len(ch_buf), cap(ch_buf))

		}
		close(ch_buf)
	}()

	time.Sleep(1 * time.Second)
	//for i := 0; i < cap(ch_buf); i++ {
	//	//fmt.Println(recover())
	//	num, ok := <-ch_buf
	//	fmt.Printf("num =%d, ok=%v\n", num, ok)
	//
	//}

	for num := range ch_buf {
		fmt.Printf("num=%d\n", num)
	}

}
