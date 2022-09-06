/*
	import:
		作用:可以导入源代码文件所依赖的package包
	规则:
		1.不能导入源代码文件中没有使用到的package,否则狗语言编译器会报编译错误
		2.import两种导入语法格式:
			2.1:
				import "package1"
				import "package2"
				import "package3"
			2.2:
				import (
					"package1"
					"package2"
					"package3"
						)

	import原理:
		1. 如果一个main导入其他包, 包将被顺序导入
		2. 如果导入的包中依赖其他包(包B),会首先导入包B, 然后初始化B包中的常量和变量, 最后如果B包中有init, 会自动执行
           init()
		3. 所有包导入完成后才会对main中的常量和变量进行初始化, 然后执行main中的init函数(如果存在init函数),最后执行main函数
		4. 如果一个包被多次导入, 则实际只会导入一次
	import特殊用法:
		1.别名操作:
			将导入的包名命名为一个容易记忆的别名
			import 别名 "包名"
		2. ”.“操作:
			"."标识的包导入后, 调用包中的函数可以省略前缀包名
		3. "_"操作:
			导入该包, 但不导入整个包, 而是执行该包中的init函数, 无法通过包名来调用包中的其他函数, 使用"_"操作往往是为了
			注册包里的引擎, 让外部可以方便的使用
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	time.Sleep(3 * time.Second)
	fmt.Print("这是import学习部分")
}
