package main

import "fmt"

/*
   range使用:
        Go语言中, range关键字用与for循环中迭代数组(array), 切片(slice), 通道(channel)或集合(map)的元素
        在数组和切片中: 返回元素的索引和索引对应的值
            for index, value := range string {
            }
        在集合中: 返回对应的key-value(键值对)对
            for key, value := range map {
            }
            如上代码中的key和value还可以省略, 假如只想读取key
                for key := range map {}
                或 for key, _ := range map {}
                如果只想读取value:
                    for _, value := range map {}
*/

func main() {
	// 定义字符串, 使用range访问字符串每个字符
	for index, value := range "我叫郭鹏涛" {
		fmt.Printf("索引%d ", index)
		fmt.Printf("数据是 %c\n", value)
	}
}
