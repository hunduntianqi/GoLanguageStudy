/*
	Ticker:
		是一个定时触发的计时器, 会以一定会的时间间隔向channel中写入当前时间, 可以以固定的时间间隔从channel中读取数据
		创建定时器:
			ticker := time.NewTicker(指定间隔时间) // 定时器每隔一定时间向channel中写入一次数据(当前时间)
		停止定时器:
			Ticker.Stop(): 定时器停止后, 不会再循环向channel中写入数据
		重置定时器:
			Ticker.Reset(指定时间): 修改原定时器指定时间, 原定时器指定时间失效
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建定时器, 指定时间间隔为1s, 定时器每隔1s向channel中写入一次数据(当前时间)
	ticker := time.NewTicker(1 * time.Second)
	ticker.Reset(2 * time.Second) // 重置定时器时间间隔为2s

	i := 0
	go func() {
		for {
			fmt.Println(time.Now())
			<-ticker.C
			i++
			fmt.Printf("i = %d\n", i)

			if i == 5 {
				ticker.Stop() // 停止定时器
			}
		}
	}()

	for {
		if i == 5 {
			break
		}
	} // 死循环防止主协程停止
}
