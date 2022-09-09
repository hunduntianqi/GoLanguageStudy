package main

import "fmt"

/*
   数组:
        是具有相同且唯一类型的一组长度固定的数据项序列, 数组类型可以是任意的原始类型(整型, 字符串等) 或自定义类型
        数组元素可以通过索引来访问或修改, 索引从0开始
        数组声明:
            var variable_name [size] variable_type
            variable_name: 数组名称
            size: 数组长度, 数组声明时必须定义数组长度
            variable_type: 数组中元素的数据类型
            例:
                var balance [10] float32 // 定义数组名为 balance, 长度为10, 元素数据类型为 float32 的数组
        数组初始化:
            格式: var 数组名 = [数组长度]数组数据类型{值1, 值2, 值3, ...}
                 或: 数组名 := [数组长度]数组数据类型{值1, 值2, 值3, ...}
            如果数组长度不确定, 可以使用"..."来代替, 编译器会根据元素个数自行推断数组长度
                例:
                    var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
                    或 balance := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
            如果设置了数组长度, 还可以通过下标来指定初始化元素:
                数组名 := [size]数组类型{索引1 : 索引1的值, 索引2:索引2的值,...索引n : 索引n的值}
        访问数组元素:
            数组通过索引来访问数组元素: 数组名[索引值]
    多维数组:
        Go语言支持多维数组, 定义方式如下:
            var variable_name [size1][size2]...[sizen] variable_type
            例:
                三维数组: var threedim [5][10][4]int
        二维数组:
            var arrayName [x][y] variable_type
            二维数组是最常见的多维数组, 二维数组可以看做一个表格, x为行, y为列
            二维数组元素访问: arrayName[行索引][列索引]
            二维数组 = append(二维数组, 一维数组): 此方法可以将一维数组添加到空的二维数组中


*/

func main() {
	var array [10]int // 定义长度为10的整型数组
	// 数组元素初始化
	for i := 0; i < len(array); i++ {
		array[i] = i + 100
	}
	// 调用函数打印数组
	arrayPrint(array)
	// 二维数组定义
	var arrayerwei [3][4]int
	// 二维数组初始化
	for i := 0; i < len(arrayerwei); i++ {
		for j := 0; j < 4; j++ {
			arrayerwei[i][j] = (i + 1) * j
		}
	}
	fmt.Println(arrayerwei) // 打印二维数组
}

// 定义函数打印数组
func arrayPrint(array [10]int) {
	// 访问数组元素
	for i := 0; i < len(array); i++ {
		if i == 0 {
			fmt.Print("{", array[i], ", ")
		} else if i == len(array)-1 {
			fmt.Println(array[i], "}")
		} else {
			fmt.Print(array[i], ", ")
		}
	}
}
