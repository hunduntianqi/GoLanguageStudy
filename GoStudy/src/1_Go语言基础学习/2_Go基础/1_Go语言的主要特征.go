package main

/*
   Go语言的主要特征:
       1. 自动垃圾回收
       2. 更丰富的内置类型
       3. 函数多返回值
       4. 错误处理
       5. 匿名函数和闭包
       6. 类型和接口
       7. 并发编程
       8. 反射
       9. 语言交互性
   Golang文件名: 所有的go源码都是以 ".go" 结尾
	Go语言命名:
		Go 的函数, 变量, 常量, 自定义类型, 包(package) 的命名方式遵循以下规则:
			1. 首字符可以是任意的Unicode字符或者下划线
			2. 剩余字符可以是Unicode字符,下划线,数字
			3. 字符长度不限
		Go 语言25个关键字:
			break default func interface select
			case defer go map struct
			chan else goto package switch
			const fallthrough if range type
			continue for import return var
		Go 语言37个保留字:
			Constants: true false iota nil
			Types: int int8 int16 int32 int64
				   uint uint8 uint16 uint32 uint64 uintptr
				   float32 float64 complex128 complex64
				   bool byte rune string error
			Functions: make len cap new append copy close delete
					   complex real imag
					   panic recover
		可见性:
			1. 声明在函数内部, 是函数的本地值, 类似于private
			2. 声明在函数外部, 对当前包可见(包内所有的.go文件都可见), 类似protect
			3. 声明在函数外部且首字母大写是所有包可见的全局值, 类似于public
		Go语言四种声明方式:
			1. var: 声明变量
			2. const: 声明常量
			3. type: 声明类型
			4. func: 声明函数
*/

func main() {

}
