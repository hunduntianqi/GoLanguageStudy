/*
	runtime包:
		1. runtime.Gosched():
			让出当前goroutine的执行权限, 调度器安排其他等待的任务执行, 并在下次某个时刻从该位置恢复执行
		2. runtime.Goexit(): 中止所在协程的运行
		3. runtime.GOMAXPROCS(): 可以设置并行计算的CPU核数最大数量, 并返回之前的值
*/

package main

import (
	"fmt"
	"runtime"
)

func task() {
	for i := 0; i < 10; i++ {
		fmt.Println("子协程i =", i)
	}
}

func main() {
	go task()
	for i := 0; i < 10; i++ {
		fmt.Println("主协程i =", i)
		// 调用runtime.Gosched()函数让出执行权限
		runtime.Gosched()
	}

}
