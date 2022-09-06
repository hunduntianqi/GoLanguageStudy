/*
	类型查询:
		interface定义为空接口时可以存储任意类型的数据(该类型实现了interface), 可以通过两种方法查询数据类型
		1. comma-ok断言
			使用if语句判断数据类型
				第一个返回值value为变量的值, 第二个返回值ok为判断结果的真假
				if value, ok := data.(预估的数据类型); ok == true {
				}
		2. switch测试
			switch value := data.(type){
				case 数据类型1:
					data为数据类型1要执行的代码
				case 数据类型2:
					data为数据类型2要执行的代码
				.....
				default:
					当所有条件都不满足时要执行的代码
			}
*/
package main

import (
	"fmt"
	"time"
)

// 定义结构体
type Studen struct {
	id   int
	name string
}

func main() {
	// 定义一个切片
	list := make([]interface{}, 3)
	list[0] = 1
	list[1] = "郭鹏涛"
	list[2] = Studen{
		id:   2,
		name: "陈欣妮",
	}

	// 类型查询 / 类型断言
	// 返回数据index为list对应下标, data为下标索引对应的值
	for index, data := range list {
		//// 使用if语句判断数据类型
		//// 第一个返回值value为变量的值, 第二个返回值ok为判断结果的真假
		//if value, ok := data.(int); ok == true {
		//	fmt.Printf("list[%d]数据类型为int, 内容为%d\n", index, value)
		//} else if value, ok := data.(string); ok == true {
		//	fmt.Printf("list[%d]数据类型为string, 内容为%s\n", index, value)
		//} else if value, ok := data.(Student); ok == true {
		//	fmt.Printf("list[%d]数据类型为Studen结构体类型, 内容为id=%d, name=%s\n", index, value.id, value.name)
		//}
		// 使用switch语句判断数据类型
		switch value := data.(type) {
		case int:
			fmt.Printf("list[%d]数据类型为int, 内容为%d\n", index, value)
		case string:
			fmt.Printf("list[%d]数据类型为string, 内容为%s\n", index, value)
		case Studen:
			fmt.Printf("list[%d]数据类型为Studen结构体类型, 内容为id=%d, name=%s\n", index, value.id, value.name)
		}
		time.Sleep(3 * time.Second)
	}

}
