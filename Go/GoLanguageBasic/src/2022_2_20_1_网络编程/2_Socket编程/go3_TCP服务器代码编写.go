/*
	网络包:net包
*/
package main

import (
	"fmt"
	"net"
)

func main() {
	// 监听
	// network: 网络类型(udp / tcp), address:IP地址和端口, 本地地址可省略不写
	listen, err := net.Listen(`tcp`, `127.0.0.1:8000`)
	if err != nil {
		fmt.Println("err =", err)
		return
	}
	defer listen.Close() // 在程序结束时关闭监听器对象
	// 阻塞等待用户链接

	conn, err := listen.Accept()
	if err != nil {
		fmt.Println("err =", err)
		return
	}
	// 连接成功后, 接受用户请求

	buf := make([]byte, 1024)    // 设置102字节的缓冲区
	n, err_buf := conn.Read(buf) // n为接受数据的长度
	if err_buf != nil {
		fmt.Println("err =", err)
		return
	}
	// 打印收到的用户请求数据
	fmt.Println("buf =", string(buf[:n]))

	defer conn.Close() // 关闭客户端链接

}
