/*
	TCP客户端
*/
package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	// 链接服务器
	// network: 网络类型(udp / tcp), address:服务器IP地址和端口
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	for {
		go func(conn2 net.Conn) {
			var send string
			//fmt.Print("请输入要发送的内容:")
			fmt.Scan(&send)
			if err != nil {
				fmt.Println("os.Stdin err =", err)
				return
			}
			// 向服务器发送键盘输入内容
			conn.Write([]byte(send))
		}(conn)
		// 接收服务器返回内容
		buf := make([]byte, 2048)
		n_read, err_read := conn.Read(buf)
		if err != nil {
			fmt.Println("read err=", err_read)
			return
		}
		if n_read == 0 || strings.Contains(string(buf), "exit") {
			fmt.Println("与服务器断开连接...")
			return
		}
		fmt.Print(string(buf[:n_read]))

	}
}
