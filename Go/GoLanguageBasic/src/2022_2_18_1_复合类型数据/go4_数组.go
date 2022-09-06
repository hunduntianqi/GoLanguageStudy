/*
	数组:
		指同一个类型的集合
		数组定义:
			var 数组名 [数组大小(元素个数)]数组类型
			注意:定义数组时, 指定数组元素个数必须为常量
		数组初始化: 数组定义同时赋值
			var 数组名 [数组大小]数组类型 =[数组大小]数组类型{元素1的值, 元素2的值, ...}
			自动推导类型初始化
				数组名 := [数组大小]数组类型{元素1的值, 元素2的值, ...}
			注意:在数组初始化时未赋值元素默认赋值为0
		二维数组:
			var 数组名 [一维数组大小][二维数组大小]数组类型
			注意:定义数组时, 有多少个"[]"就有几维数组
		数组的比较:
			比较数组每一个元素是否一样
			数组元素一样:true
			数组元素不一样:false
		同类型数组可以赋值:
			数组1 = 数组2
		数组做函数参数, 为值传递
			格式:
				func 函数名(参数名 [数组大小]数组类型) {
					功能代码
				}
		数组指针做函数参数:
			格式:
				func 函数名(参数名 *[数组大小]数组类型) {
					功能代码
				}
*/

package main

import (
	"fmt"
	"go/types"
)

type ArrFunc func(arr types.Slice)

// 冒泡排序
func sortArr(arrInt [10]int) [10]int {
	// 冒泡排序
	for j := 0; j < len(arrInt)-1; j++ {
		for i := 0; i < len(arrInt)-1; i++ {
			if arrInt[i] < arrInt[i+1] {
				arrInt[i], arrInt[i+1] = arrInt[i+1], arrInt[i]
			}
		}
	}
	return arrInt
}
func main() {
	//定义整型数组
	var arrInt [10]int
	for i := 0; i < len(arrInt); i++ {
		arrInt[i] = i * i
		fmt.Println(arrInt[i])
	}
	//
	//for i, data := range arrInt {
	//	fmt.Printf("%d:%d\n", i, data)
	//}

	//var arrInt2 [10][10]int
	//for i := 1; i < len(arrInt2); i++ {
	//	for j := 1; j <= i; j++ {
	//		arrInt2[i][j] = i * j
	//		fmt.Printf("%d*%d=%d\t", j, i, arrInt2[i][j])
	//	}
	//	fmt.Println()
	//	//fmt.Println(arrInt2[i])
	//}
	arrInt = sortArr(arrInt)
	fmt.Println(arrInt)
}
