package main

import "fmt"

func main() {
	str := "abc"
	// 使用range迭代打印字符串元素
	// range默认返回两个参数:1.元素的下标索引, 2. 元素本身, 因此需要使用两个变量类接受range返回结果
	for i, daat := range str {
		fmt.Printf("str[%d] = %c\n", i, daat)
	}
}
