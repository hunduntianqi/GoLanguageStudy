package main

import (
	"fmt"
	"net"
	"strings"
)

type Client struct {
	C    chan string // 用于发送数据的管道
	Name string      // 用户名
	Addr string      // 网络地址
}

// 保存在线用户, 通过IP和端口
var online map[string]Client

var message = make(chan string)

// 新开一个协程, 转发消息, 当有消息发送时, 遍历map, 给每个成员发送消息
func Messager() {
	// 给map分配空间
	online = make(map[string]Client)
	for {
		msg := <-message // 没有消息前, 此处会被阻塞
		// 遍历map, 给每个成员发送消息
		for _, cli := range online {
			cli.C <- msg
		}
	}
}

func WriteMsgToClient(cli Client, conn net.Conn) {
	for msg := range cli.C {
		conn.Write([]byte(msg + "\n"))
	}
}

func MakeMsg(cli Client, msg string) string {
	return `[` + cli.Addr + `]` + cli.Name + msg
}

// 处理用户连接
func HandleConn(conn net.Conn) {
	defer conn.Close()
	// 获取客户端的网络地址
	cliAddr := conn.RemoteAddr().String()
	// 创建结构体,默认, 用户名和网络地址一样
	cli := Client{C: make(chan string), Name: cliAddr, Addr: cliAddr}
	// 将结构体加入map
	online[cliAddr] = cli
	// 广播某人在线
	fmt.Println(MakeMsg(cli, `: login`))
	message <- MakeMsg(cli, ": login")
	go WriteMsgToClient(cli, conn)
	//新建协程, 接受用户发送过来的数据
	func() {
		for {
			buf := make([]byte, 2048)
			n, err := conn.Read(buf)
			if n == 0 { // 对方断开或有异常
				fmt.Println("err =", err)
				// 将当前连接对象关闭并从map中剔除
				delete(online, cliAddr)
				return
			}

			//WriteMsgToClient(cli, conn)}()
			// 查看当前用户
			if strings.Contains(string(buf[:n]), "用户") {
				//遍历map, 像当前用户发送在线人员名单
				conn.Write([]byte("当前在线人员为:\n"))
				for _, value := range online {
					conn.Write([]byte(value.Addr + value.Name + "\n"))
				}
			} else if strings.Contains(string(buf), "rename /") {
				cli.Name = strings.Split(string(buf), "/")[1]

			} else {
				// 转发消息
				message <- MakeMsg(cli, ":"+string(buf[:n]))
			}
		}
	}()
	for {

	}
}

func main() {
	fmt.Println("服务器启动, 等待客户端连接...")
	// 监听
	listen, listn_err := net.Listen(`tcp`, `127.0.0.1:8000`)
	if listn_err != nil {
		fmt.Println(`listn_err =`, listn_err)
		return
	}
	defer listen.Close()

	// 新开一个协程, 转发消息, 当有消息发送时, 遍历map, 给每个成员发送消息
	go Messager()

	// 主协程, 循环阻塞等待用户连接
	for {
		conn, conn_err := listen.Accept()
		if conn_err != nil {
			fmt.Println("conn_err =", conn_err)
			continue
		}
		// 处理用户请求
		go HandleConn(conn)
	}

}
