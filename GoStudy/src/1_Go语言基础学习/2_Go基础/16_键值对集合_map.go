package main

import "fmt"

/*
   map:
       map是一种基于key-value的数据结构, go语言中的map是引用类型, 必须初始化才能使用, 如果不初始化, 就会创建一个nil map, nil map
       不能用来存放键值对, go语言中的Map是一种无序的集合, 也可以向迭代数组和切片一样迭代, 但无法决定返回顺序
        map定义:
            var map_name map[key_type]value_type
        或使用make关键字定义map
            map_name := make(map[key_type]value_type)
            map_name: 定义Map的名称
            key_type: map中键的数据类型
            value_type: map中值的数据类型
        Map集合初始化: map_name = make(map[key_type]value_type)In
        为Map集合添加元素:
            map_name[key_name] = value
            注意: 集合中键名不能重复, 重复键名重复时, 会覆盖原来的值
        删除Map集合元素:
            delete(map_name, key_name): 删除集合中对应键名的元素
        判断key是否存在:
			value, bool := map名[key]
			如果key存在, 则bool值为true, value为对应键的值, key不存在, bool值为false
        map做函数参数-引用传递:
			func 函数名(参数名 map[键类型]值类型) {
			}
*/

func main() {
	// 定义Map集合
	var countyCapitalMap map[string]string     // 创建国家首都Map集合
	countyCapitalMap = make(map[string]string) // Map集合初始化
	// map集合添加元素
	countyCapitalMap["France"] = "巴黎"
	countyCapitalMap["Italy"] = "罗马"
	countyCapitalMap["Japan"] = "东京"
	countyCapitalMap["India"] = "新德里"
	fmt.Println(countyCapitalMap)
	// for循环遍历Map集合
	for key, value := range countyCapitalMap {
		fmt.Printf("%s的首都是%s\n", key, value)
	}
	// countyCapitalMap["France"] = "伦敦" // map集合添加键名重复元素时, 会覆盖原来的值
	// 删除元素
	delete(countyCapitalMap, "France") // 删除键名为"France"的元素
	fmt.Println(countyCapitalMap)
	countyCapitalMap["Chinese"] = "北京"
	fmt.Println(countyCapitalMap)
	// 判断Map中是否存在某个键名
	value, ok := countyCapitalMap["Chinese"]
	if ok {
		fmt.Println("Chinese 的首都为:", value)
	} else {
		fmt.Println("countyCapitalMap集合中无Chinese相关的数据")
	}
	printMap(countyCapitalMap) // 在函数中遍历Map后删除Japan元素
	fmt.Println(countyCapitalMap)
}

// 定义函数遍历Map集合
func printMap(mapDemo map[string]string) {
	for key, value := range mapDemo {
		fmt.Printf("mapDemo[%s] = %s\n", key, value)
	}
	// map用作函数参数为引用传递, 在函数内部操作集合会对外部集合产生影响
	delete(mapDemo, "Japan")
}
