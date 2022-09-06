/*
	关键字:
		Go中的保留关键字只有25个:
			break,default,func,interface,select,case,defer,go,map,struct,chan,else,goto,package,switch
			const,fallthrough, if,range,type,continue,for,import,return,var
	Go语言中的36个预定标识符, 包括基础数据类型和系统内嵌函数:
		append,bool,byte,cap,close,complex,complex64,complex128, unit16, copy,false,float32,float64,
		imag,int,int8,int16,uint32,int32,int64,iota,len,make,new,nil,panic,uint64,print,println,real,
		recover, string,true,unit,uint8,uintprt
	注释形式:
		单行注释://
		多行注释:/*...
				...*/

/*
	全局变量:
		main函数外定义的变量
*/
// 程序所属的包, 文件代码开头第一行必须有对应的包
package main

// 导入依赖包
import "fmt"

// 常量定义
const NAME string = "在学习Go语言"

// 全局变量的声明与赋值, 在main函数外定义的变量
var a string = "郭鹏涛"

// 一般类型的声明
type study string

// 结构的声明
type Learn struct {
}

// 声明接口
type LearnInter interface {
}

// 函数定义
func learngo() {
	fmt.Print("我是郭鹏涛\n", a, NAME)
}

// main函数, 程序入口
func main() {
	learngo()
}
