package main

import "fmt"

/*
   取地址符:
       GO语言取地址符"&"，放到变量前使用可以返回该变量的内存地址
   指针:
        一个指针变量指向了一个值得内存地址
        声明格式:
            var var_name *var_type
            var_name: 指针变量的名称
            var_type: 指针类型
            *: 代表定义的变量是作为一个指针使用
        指针使用流程:
            1. 定义指针变量
            2. 为指针变量赋值 (指针代表一个变量的内存地址, 在给指针变量赋值时, 实际是将变量的内存地址赋值给指针)
                指针变量名 = &要赋值地址给指针的变量名
            3. 访问指针变量中指向地址的值
                在指针类型前面加上"*"号获取指针所指向的内容
                即: *指针变量名 == 要赋值地址给指针的变量名
                注意: 由于指针直接操作的是变量的内存地址, 当地址发生变化时, 指针存储的值也会发生变化, *指针变量名的值也会
					 发生相应的改变
		空指针:
			当一个指针变量定义后没有赋值, 该指针就是一个空指针, 空指针的值为nil(相当于null), 表示零值或空值
		指针数组:
			定义格式:
				var array_name [size] *var_type
				array_name: 指针数组名
				size: 指针数组长度
				var_type: 指针数组类型
				*指代这是一个指针数组
			与普通数组相比, 指针数组中存储的元素都是指针
		指向指针的指针:
			如果一个指针变量存放的是另一个指针变量的地址, 这个指针变量为指向指针的指针变量, 定义格式如下:
				var point_name **var_type
				**: 代表指向指针的指针
				指向指针的指针中存放的值是指针地址, 使用"*"变量名的方式取出的是指针存储的普通变量的地址, "**变量名"
				才可以最终取出普通变量的值
				即:
					普通变量a, 指针变量b, 指向指针的指针变量c
					则:
						a = *b = **c  // 变量a存储的实际值
						&a = b = *c  // 变量a的内存地址值
						c = &b  // 指针b的内存地址值
		指针作为函数参数:
			Go语言中指针可以作为函数参数, 只需要将函数参数设置为指针形式即可
				func 函数名 (指针形式的形参) (返回值类型列表) {
					函数体
				}
			指针存储的是变量地址, 因此指针作为参数传递, 在函数内部的操作, 会改变指针指向的外部变量的值
*/

func main() {
	// 定义变量并赋值
	var num int = 10
	// 打印变量内存地址
	fmt.Printf("变量num的内存地址是:%p\n", &num)
	// 定义指针变量
	var point *int
	// 指针变量赋值
	point = &num // 指针代表一个变量的内存地址, 在给指针变量赋值时, 实际是将变量的内存地址赋值给指针
	fmt.Printf("point的值是:%v\n", point)
	fmt.Printf("变量num的值是:%d\n", num)
	fmt.Printf("*point的值是:%d\n", *point)
	// 调用方法, 打印指针数组执行结果
	pointArray()
	pointPoint()
	// 调用函数传递指针作为形参
	fmt.Printf("num = %d\n", num)
	num2 := pointParam(point) // 调用函数后, 指针指向的外部变量num同步被修改
	fmt.Printf("num = %d, num2 = %d\n", num, num2)
}

// 定义指针数组并赋值
func pointArray() {
	// 定义普通整形数组
	var array = [3]int{1, 2, 3}
	// 定义指针数组
	var arrayPoint [3]*int
	// 指针数组初始化
	for i := 0; i < len(arrayPoint); i++ {
		arrayPoint[i] = &array[i]
	}
	// 通过*指针访问普通数组元素
	for i := 0; i < len(arrayPoint); i++ {
		fmt.Printf("arrayPoint[%d] = %p, *arrayPoint[%d] = %d, array[%d] = %d\n", i, arrayPoint[i], i, *arrayPoint[i],
			i, array[i])
	}
}

// 定义函数演示指向指针的指针
func pointPoint() {
	// 定义普通整型变量
	var num int = 10
	// 定义指针变量
	var num_Point *int = &num
	// 定义指向指针的指针变量
	var num_Point_Point **int = &num_Point
	// 指向普通变量地址
	fmt.Printf("*num_Point_Point = %p, num_Point = %p, &a = %p\n", *num_Point_Point, num_Point, &num)
	// 指向指针变量地址
	fmt.Printf("num_Point_Point = %p, &num_Point = %p\n", num_Point_Point, &num_Point)
	// 指向普通变量存储的实际值
	fmt.Printf("num = %d, *num_Point = %d, **num_Point_Point = %d\n", num, *num_Point, **num_Point_Point)
}

// 定义函数将指针作为参数
func pointParam(point *int) int {
	// 修改指针指向变量的值, 并返回
	*point = 100
	return *point
}
