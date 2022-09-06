/*
	并行(parallel):
		指在同一时刻, 有多条指令在多个处理器上同时执行
	并发(concurrency):
		指在同一时刻只能有一条指令执行, 但多个进程指令被快速的轮换执行, 使得在宏观上具有多个进程同时执行的
	效果, 但在微观上并不是同时执行的
	Go语言并发优势:
		Go语言从语言层面就支持了并发, 同时, Go语言提供类自动垃圾回收机制, 可以
		方便的进行并发程序的内存管理
		Go语言为并发编程内置了上层API基于CSP(communicating sequential processes, 顺序通信进程)模型, 这
	意味着显示锁都是可以避免的, 因为Go语言通过安全的通道发送和接受数据以实现同步, 大大简化了并发编程的编写

*/
package main

func main() {

}
