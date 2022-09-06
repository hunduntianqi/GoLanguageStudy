/*
	切片(slice):
		通过内部指针和相关属性引用数组片段, 来实现大小不固定, slice是一个引用类型数据
		定义格式1:
			var 切片名 [] 数据类型
			切片定义与数组类似, 只是不需要指定元素个数, 指定元素个数为数组, 不指定元素个数为切片
		定义格式2:
			var 切片名 []数据类型 = make([]数据类型, 切片长度)
		获取切片容量:cap(切片名称)
		给切片末尾追加元素:append(切片名称, 追加元素)
			append扩容特点:
				append函数会智能检测底层数组容量的增长,当超过原底层数组的容量,一般会以二倍容量重新分配底层数组
			并复制原来的数据
		copy函数:copy(切片1, 切片2)
			将切片2的元素复制并覆盖到切片1对应的位置
		切片做函数参数-引用传递:
			格式:
				func 函数名(参数名 []数据类型) {
					功能代码
				}


*/

package main

func main() {

}
