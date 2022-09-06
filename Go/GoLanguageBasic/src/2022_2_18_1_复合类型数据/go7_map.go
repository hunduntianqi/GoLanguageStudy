/*
	Go语言中的map(映射 / 字典):
		是一种内置的数据结构,是一个无序的key-value对的集合
		在一个map中的键都是唯一的,而且键不能为具有引用语义的数据类型, 切片, 函数以及包含切片的结构类型不能作为键
		map定义格式:
			var map名 map[键类型]值类型
			注意: 该方式定义map中添加元素前必须先初始化
		make创建map:
			map名 := make(map[键类型]值类型)
		map赋值:
			map名[key] = value
			如果key已存在, 会修改原来key的内容
		map遍历:
			for key value := range map名 {

			}
		map删除元素:
			delete(map名, key)
		判断key是否存在:
			value, bool := map名[key]
			如果key存在, 则bool值为true, key不存在, bool值为false
		map做函数参数-引用传递:
			func 函数名(参数名 map[键类型]值类型) {
			}
*/

package main

import "fmt"

func main() {
	map2 := make(map[int]string)
	// map添加元素
	map2[1] = "郭鹏涛"
	map2[2] = "陈欣妮"
	map2[3] = "郭鹏涛Love陈欣妮"
	fmt.Println(map2)

}
