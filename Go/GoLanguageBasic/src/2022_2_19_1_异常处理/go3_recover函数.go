/*
	recover函数:
		Go语言提供的专门用于拦截运行时产生panic异常的内建函数, 可以使当前的程序从运行时panic的状态中恢复并重新
	获得流程控制权
		recover函数:
			func recover() interface {}
		注意: recover函数只有在defer修饰的函数中有效
		recover使用原理:
			当定义defer语句的函数发生panic异常时, 如果调用了recover函数, 该函数会使程序从panic中恢复过来, 并返回panic value,
		导致panic异常的函数不会再继续运行, 但能正常返回; 在未发生panic时调用recover, recover函数会返回nil
*/

package main

import "fmt"

func recover_test(x int) {
	// 定义匿名函数, 使用recover函数处理异常, 避免程序崩溃
	defer func() {
		// 判断是否产生异常, 产生异常打印异常信息
		if err := recover(); err != nil {
			fmt.Println(err) // 打印异常信息
		}
	}()
	var a [10]int
	a[x] = 10
	fmt.Printf("a[%d]=%d\n", x, a[x])
}

func main() {
	recover_test(3)
	recover_test(5)
	recover_test(11)
	recover_test(10)
	recover_test(9)

}
