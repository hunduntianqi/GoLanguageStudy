/*
	Append系列函数:将整数等转换为字符串后, 添加到现有的字节数组中
	Format系列函数:将其他类型数据转换为字符串
	Parse系列函数:把字符串转换为其他类型
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 定义字节切片
	list := make([]byte, 0, 1024)
	list = strconv.AppendBool(list, false)
	// 方法中第二个参数为要追加的数据, 第三个参数为指定int进制类型, 下列为以10进制追加数据
	list = strconv.AppendInt(list, 100, 10)
	list = strconv.AppendQuote(list, "郭鹏涛")
	fmt.Printf("list = %v \n", string(list)) // 需要转换为字符串再打印,否则会以ascll编码数字显示
	// 其他数据类型转换为字符串
	var str string
	str = strconv.FormatBool(false)
	fmt.Printf("str=%s\n", str)
	// 'f'指打印格式, 以小数方式, -1指小数点后位数, 64指以float64位处理
	str = strconv.FormatFloat(3.1415926, 'f', -1, 64)
	fmt.Printf("str=%s\n", str)
	// 整型转字符串
	str = strconv.Itoa(123) // 该方式为整型转字符串最常用的方式
	fmt.Printf("str=%s\n", str)

	// 字符串转其他数据类型
	bool, err_bool := strconv.ParseBool("true")
	if err_bool == nil {
		fmt.Printf("bool=%v\n", bool)
	}
	num, err_int := strconv.ParseInt("123", 10, 64)
	if err_int == nil {
		fmt.Printf("num=%v\n", num)
	}
}
