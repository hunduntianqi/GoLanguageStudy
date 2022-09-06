package main

import (
	"fmt"
	"net"
	"strings"
)

// 定义函数处理用户请求
func HandleConn(conn net.Conn) {
	// 函数调用完毕, 关闭conn
	defer conn.Close()
	// 获取客户端网络地址信息
	addr := conn.RemoteAddr().String()
	fmt.Println(addr, " connect sucessful")
	buf := make([]byte, 2048)
	for {
		// 读取用户数据
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("err=", err)
			return
		}
		if string(buf[:n]) == "exit" {
			fmt.Println(addr, "已断开连接")
			conn.Close()
			return
		}
		fmt.Printf("%v:read buf=%s", addr, string(buf[:n]))

		// 将数据转换为大写, 再给用户发送
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}
func main() {
	fmt.Println("服务器已启动, 等待客户端链接...")
	// 监听
	linstn, err := net.Listen(`tcp`, `127.0.0.1:8000`)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer linstn.Close()
	// 接收多个用户请求
	for {
		conn, err := linstn.Accept()
		if err != nil {
			fmt.Println("err=", err)
			return
		}

		// 并发处理用户请求,新建子协程
		go HandleConn(conn)

	}

}
