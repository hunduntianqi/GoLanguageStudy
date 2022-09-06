/*
	package:
		1. 是go语言中最基本的分发单位和工程管理中依赖关系的体现
		2. 每个gu语言源码文件开头都必须有一个package声明,表示源码文件所属代码包
		3. 要生成go语言可执行文件, 必须要有main的package包, 且必须在该包下有main()函数
		4. 同一路径下只能存在一个package,一个package可以拆成多个源文件组成


*/

package main

import "fmt"

func main() {
	fmt.Print("你好, 这是package的详解内容")
}
