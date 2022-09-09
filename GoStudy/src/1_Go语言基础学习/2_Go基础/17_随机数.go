package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
   Go语言中的随机数是伪随机数, math/rand包实现了伪随机数生成器
   Go语言中随机数生成需要设置种子, 如果不设置种子, 随机数每次运行结果相同, 默认种子为1, 相同种子产生的随机数是相通的
   设置种子:
       为了保证种子不是固定的, 使用time这个包来调取当前时间, 采用当前时间的纳秒作为种子来生成随机数
       格式:
            rand.Seed(time.Now().Unix)
            randomNumber := rand.数据类型(max_num)
            根据数据类型和max_num来随机生成随机数, 随机数范围为[0, max_num), 左闭右开
            rand.Float32() 和 rand.Float64() 默认生成的随机数范围为[0.0, 1.0), 左闭右开, 通过算术运算符可以获得其他范围的随机数
            rand.Int31n(num) 和 rand.Int63n(num) 可以随机生成[0, num)范围内的整数, 左闭右开
*/

func main() {
	// 设置种子
	rand.Seed(time.Now().Unix())
	// 生成随机数
	randomNumber := rand.Int63n(1000) // 生成[0, 999)的随机数, 左闭右开
	fmt.Println(time.Now().Unix())
	fmt.Printf("生成的随机数randomNumber = %d\n", randomNumber)
}
