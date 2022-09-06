/*
	Timer:
		是一个定时器对象, 代表未来的一个单一事件
		创建Timer定时器:
			timer := time.NewTimer(指定时间)
			指定时间后会向timer通道(channel)写入当前时间: time_now := <- timer.C
		通过Timer可以实现延时功能
		定时器停止:
			Timer.stop(): 定时器停止后, 不会再等待执行写入定时器当前时间
		定时器重置:
			Timer.Reset(指定时间): 修改原定时器指定时间, 原定时器指定时间失效
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个定时器, 设置时间为2s, 2s后, 向timer通道写入当前时间
	timer := time.NewTimer(5 * time.Second)
	t_now := time.Now()
	fmt.Printf("当前时间为:%v\n", t_now)
	t := <-timer.C
	fmt.Printf("t_Gap = %v\n", t.Sub(t_now))
	//<-timer.C // 此处会阻塞死锁

}
