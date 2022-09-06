/*
	正则表达式:
		是一种进行模式匹配和文本操作的复杂而又强大的工具, 按照正则表达式的语法规则, 随需构造出的匹配规则能够从原始文本中
	筛选出几乎任何想要得到的字符组合
	Go语言通过regexp标准包为正则表达式提供了官方支持
	regexp包使用流程:
		1. 指定正则表达式匹配规则
			reg := regexp.MustComlile('正则匹配规则')
		2. 根据匹配规则提取信息
			reg.FindAllString(buf, -1):返回数据为只有一个参数的切片
			reg.FindAllStringSubmatch(buf, -1): 返回数据为多个切片组成的切片
*/

package main

import (
	"fmt"
	"regexp"
)

func main() {
	var buf string = "abc azc a7c aac 888 a9c tac"
	// 指定匹配规则
	reg, err := regexp.Compile(`a.c`)
	if err == nil {
		fmt.Println(reg)
	}
	fmt.Println(reg.FindAllString(buf, -1))
	fmt.Println(reg.FindAllStringSubmatch(buf, -1))
	for index, data := range reg.FindAllStringSubmatch(buf, -1) {
		fmt.Printf("data[%d]=%s\n", index, data[0])
	}
}
