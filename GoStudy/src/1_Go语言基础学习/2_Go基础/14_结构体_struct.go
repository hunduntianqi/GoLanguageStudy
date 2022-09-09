package main

/*
	自定义类型:
		Go语言中可以使用type关键字来自定义数据类型, 可以基于内置的基本数据类型定义, 也可以通过struct定义
		例:
			type NewInt int ==> 定义NewInt是一种新的数据类型, 具有int的特性
		类型别名设置:
		type 类型别名 = 类型名
		例: type MyInt = int ==> MyInt是int类型的别名, 本质上MyInt和int是同一种类型
   	结构体:
        Go语言中由一列具有相同的或不同的类型数据构成的集合 (个人认为类似于Java中的类, 只有成员变量, 没有成员方法)
        结构体定义:
            type struct_name struct {
                member1 datatype
                member2 datatype
                member3 datatype
                .....
            }
        结构体被定义后, 就可以像其他基本数据类型一样进行变量声明 (个人认为类似于Java中的创建对象)
            格式: variable_name := struct_name {member1_value, member2_value,...}
                或 variable_name := struct_name {key1:value1, key2:value2,...}
                或 variable_name := struct_name {key1:value1, key3:value3, ...} 未初始化成员变量为基本类型默认值
        访问结构体成员:
            格式: 结构体名.成员名称
        修改结构体成员变量的值:
            格式: 结构体名.成员名称 = new_value
		结构体作为函数参数:
			func func_name (variable_name struct_name) (return list) {
				函数体
			}
			该传递类型为值传递, 函数内部对结构体做的修改不会影响函数外部结构体的数据
		结构体指针:
			类似于其他指针变量, 定义时指针类型为结构体名称
			var struct_point *struct_name
			结构体指针赋值: struct_point = &结构体类型变量名
			使用结构体指针访问结构体变量成员: struct_point.成员字段名
*/

import "fmt"

// 定义学生结构体
type Student struct {
	name string
	age  int
	sex  rune
}

func main() {
	// 声明学生类型变量
	var student1 Student
	var student2 Student
	var student3 Student
	// 结构体成员初始化方式一
	student1 = Student{"郭鹏涛", 24, '男'}
	name := student1.name
	age := student1.age
	sex := student1.sex
	fmt.Printf("student1: name==>%s, age==>%d, sex==>%v\n", name, age, string(sex))
	// 结构体成员初始化方式二
	student2 = Student{name: "郭鹏涛", sex: '男', age: 24}
	fmt.Println("student2: ", student2.name, student2.age, string(student2.sex))
	// 结构体成员初始化方式三
	student3 = Student{name: "郭鹏涛", age: 24}
	fmt.Println("student3: ", student3.name, student3.age, student3.sex)
	// 修改结构体成员变量的值
	student2.name = "郭鹏强"
	student2.age = 22
	fmt.Println("student2: ", student2.name, student2.age, string(student2.sex))
	// 调用函数, 参数传入结构体变量
	structFunc(student3)
	fmt.Println("student3: ", student3.name, student3.age, student3.sex)
	// 调用函数, 参数传入结构体变量内存地址作为指针
	structPoint(&student1)
}

// 定义参数为结构体的函数
func structFunc(variable Student) {
	variable.sex = '男'
	fmt.Println(variable.name, variable.age, string(variable.sex))
}

// 定义参数为结构体指针的函数
func structPoint(point *Student) {
	// 通过指针访问结构体成员
	fmt.Println(point.name, point.age, string(point.sex))
}
