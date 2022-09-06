package main

import (
	"fmt"
	"math/rand" // 随机数模块
	"time"
)

// 生成一个四位随机数
func creatNum() int {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(10000)
	if num >= 1000 {
		return num
	} else {
		return -1
	}
}
func main() {
	rand.Seed(time.Now().UnixNano())
	var randNum int
	// 产生一个4位的随机数
	randNum = creatNum()
	fmt.Println("随机数randNum:", randNum)
	for i := 0; i < 3; i++ {
		fmt.Println("请输入您猜的数字:")
		var a int
		fmt.Scanf("%d\n", &a)
		if a == randNum {
			fmt.Println("恭喜您猜对了！！")
			break
		} else {
			if i == 2 {
				fmt.Println("您已猜错三次, 游戏退出~~")
			} else {
				fmt.Printf("您猜错了, 您还有%d次机会\n", 3-1-i)
			}
		}
	}
}
