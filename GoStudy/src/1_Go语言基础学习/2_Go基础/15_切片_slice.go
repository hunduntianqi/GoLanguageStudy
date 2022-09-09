package main

/*
   切片(slice):
		通过内部指针和相关属性引用数组片段, 来实现大小不固定, slice是一个引用类型数据
       定义切片:
           var identifier []type
           identifier: 切片名称
           type: 切片中存储元素的数据类型
       使用make()函数创建切片:
           var identifier []type = make([]type, len, capacity)
           或简写为: identifier := make([]type, len, capacity)
           len: 指定切片底层数组的长度, 也是切片的初始长度
           capacity: 指定切片容量
		   如果指定切片的容量, 则切片的初始长度不能大于容量
       len()函数: 可以获取当前切片的长度 (元素个数)
       cap()函数: 可以获取当前切片的容量 (最大可存储元素个数)
	   append(splice, value)函数: 向切片追加元素, 并复制原切片所有内容后返回新的切片(需要有变量来接收新的切片)
								 当追加元素时, 如果 len = cap,则切片底层数组会进行扩容操作, 从而形成一个新切片(与原切片的指针不同)
								 即与原切片不是同一个切片, 对新切片的操作不会影响到原切片(即使新切片与原切片同名)
								 切片扩容规则:
									原切片长度小于1024: 新切片容量翻倍
									原切片长度大于或等于1024: 新切片容量为原切片容量的1.25倍
									当同时添加的数据导致新切片的容量大于原切片容量的2倍时, 会使用需要的容量作为新切片的容量
										内置数据类型: len为奇数, cap = len + 1; len为偶数, cap = len
										自定义类型: cap = 旧切片容量 + 添加元素个数
		删除切片元素:
			Go语言中并未提供对切片元素的删除操作, 因此需要通过切片自己的特性来进行元素删除操作
			删除切片开头的元素:
				方式1: slice = slice[n:] ==> 删除切片开头n个元素; 该方法实际是产生了新的切片, 改变了底层数组元素指针, 底层元素
						数组指针会变为第n+1个元素的指针
				方式2: slice = append([:0], [n:]...) "..."不能省略==> 删除切片开头n个元素; 该方法实际操作的是原切片对应的内存空间,
					   不会改变切片的内存结构, 改变的是原来的切片
				方式3: slice = slice[:copy(splice, splice[n:])] ==> 删除切片开头n个元素; 该方法实际操作的是原切片对应的内存空间,
						不会改变切片的内存结构, 改变的是原来的切片
			删除切片中间的元素:
				方式1: slice = append([:n-1], [n + i:]...) "..."不能省略==> 删除切片中间i个元素; 该方法实际操作的是原切片对应的
						内存空间, 不会改变切片的内存结构, 改变的是原来的切片
				方式2: slice = slice[:n - 1 + copy(slice[n - 1:], slice[n+i:])] 删除切片中间i个元素
			删除尾部元素:
					slice = slice[:len[slice] - i]: 删除尾部i个元素, 该方法不会修改切片内存结构
			删除切片元素总结:
				在使用截取切片的方式删除开头元素时, 由于直接在内存中删除了开头部分的元素, 而不是修改元素的值, 会导致切片底层数组指针
				改变, 因此相当于新建了一个切片; 其他删除数组元素的方式均未改变数组指针, 切片的容量未发生改变, 因此数组结构未发生改变

       空切片:
           一个切片在未初始化之前默认为nil, 长度为0
		切片截取:
			可以通过设置上下限来截取切片[lower_bound:upper_bound] ==> 包左不包右
		copy()函数:
			copy(slice1, slice2): 将slice2中的元素复制到slice1中, 如果len(slice1) < len(slice2), 则slice2中的元素无法全部复制
								  到slice1中
		切片遍历:
			for index, data := range slice {
			}
			index: 切片索引值
			data: 切片索引对应的数据
		切片做函数参数-引用传递:
			格式:
				func 函数名(参数名 []数据类型) {
					功能代码
				}
			在函数中传递切片操作, 如果因追加元素发生扩容操作, 产生一个新的切片, 则在函数中对切片的操作可能不会改变外部切片的数据
			因为不是同一个切片, 底层数组指针不同
*/

import "fmt"

func main() {
	// 定义一个长度3, 容量5的切片
	list := make([]int, 5, 5)
	fmt.Println(list)
	fmt.Printf("list切片的初始长度len = %d, 容量cap = %d\n", len(list), cap(list))
	// for循环给切片元素赋值
	for i := 0; i <= 4; i++ {
		list[i] = i
	}
	list = append(list, 5, 6, 7, 8, 9, 10)
	fmt.Println(list)
	fmt.Printf("list切片添加数据后长度len = %d, 容量cap = %d\n", len(list), cap(list))

	// sliceCapCheck()
	sliceSplit()
	// 切片遍历
	for index, data := range list {
		fmt.Printf("list[%d] = %d\n", index, data)
	}
	changeSlice(list)
	fmt.Println(list)
	deleteSpliceElement()
}

// 定义函数验证切片扩容规则
func sliceCapCheck() {
	// 定义切片, 容量小于1024
	list1 := make([]int, 510, 512)
	fmt.Println("扩容前: list1 len = ", len(list1), ", cap = ", cap(list1))
	// 使用append函数向list1添加元素
	list1 = append(list1, 1, 2, 3, 4)
	fmt.Println("扩容后: list1 len = ", len(list1), ", cap = ", cap(list1))
}

// 定义函数截取切片
func sliceSplit() {
	// 定义切片
	var list []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// 截取切片
	fmt.Println(list[2:5]) // 包左不包右, 截取切片索引2-4
	fmt.Println(list[2:])  // 从索引2开始截取切片
	fmt.Println(list[:5])  // 从索引0开始截取切片到索引4
	var list2 = make([]int, 4, 4)
	copy(list2, list)
	fmt.Println(list2)
}

// 切片做函数参数==>引用传递
func changeSlice(list []int) {
	// 给切片添加元素
	list = append(list, 19980224)
	fmt.Println(list)
	// 通过append函数 删除切片第5个元素
	list = append(list[0:4], list[5:]...)
}

// 定义函数观察切片删除元素的内存变化
func deleteSpliceElement() {
	// 定义切片
	var list []int32 = []int32{1, 2, 3, 4, 5}
	fmt.Println(list)
	fmt.Printf("list ==> len = %d, cap = %d\n", len(list), cap(list))
	fmt.Println("list的内存地址为:", &list[0])
	// 使用切片方式删除第一个元素
	list = list[1:] // 产型新的数组, 会产生一个新的切片, 与原切片指针不同
	fmt.Println(list)
	fmt.Printf("list ==> len = %d, cap = %d\n", len(list), cap(list))
	fmt.Println("list的内存地址为:", &list[0])
	// 使用append函数删除第一个元素
	list = append(list[:0], list[1:]...) // 不会改变原来切片的内存结构
	fmt.Println(list)
	fmt.Printf("list ==> len = %d, cap = %d\n", len(list), cap(list))
	fmt.Println("list的内存地址为:", &list[0])
	// 使用copy删除第一个元素
	list = list[:copy(list, list[1:])] // 不会改变原来切片的内存结构
	fmt.Println(list)
	fmt.Printf("list ==> len = %d, cap = %d\n", len(list), cap(list))
	fmt.Println("list的内存地址为:", &list[0])
	list = append(list, 1, 2, 3, 5, 4)
	fmt.Println(list)
	fmt.Printf("list ==> len = %d, cap = %d\n", len(list), cap(list))
	fmt.Println("list的内存地址为:", &list[0])
	// 使用append删除第4个元素
	list = append(list[:3], list[4:]...) // 不会改变原来切片的内存结构
	fmt.Println(list)
	fmt.Printf("list ==> len = %d, cap = %d\n", len(list), cap(list))
	fmt.Println("list的内存地址为:", &list[0])
	// 使用copy删除第4个元素
	list = list[:2+copy(list[3:], list[4:])]
	fmt.Println(list)
	fmt.Printf("list ==> len = %d, cap = %d\n", len(list), cap(list))
	fmt.Println("list的内存地址为:", &list[0])
	copy(list, []int32{1, 2, 3, 4, 5})
	// 删除尾部两个元素
	list = list[:3]
	fmt.Println(list)
	fmt.Printf("list ==> len = %d, cap = %d\n", len(list), cap(list))
	fmt.Println("list的内存地址为:", &list[0])
}
