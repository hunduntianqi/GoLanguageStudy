/*
	new函数的使用:
		表达式new(T)可以创建一个T类型的匿名变量,为T类型的新值分配并清零一块内存空间, 然后将内存空间的地址作为结果返回,
	返回的结果为指向新的T类型值的指针值, 返回的指针类型为*T
*/

package main

import "fmt"

func main() {
	var p1 *int
	p1 = new(int)
	fmt.Printf("p1=%p\n", p1)

	*p1 = 100
	fmt.Printf("p1=%d\n", *p1)

}
