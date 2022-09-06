package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

// 定义函数处理保存接受的文件
func save_file(file_data []byte) {
	// 调用匿名函数
	defer func() {
		// 判断是否产生异常
		if err := recover(); err != nil {
			fmt.Println(err)
			return
		}
	}()
	// 通过split函数切割file_data, 提取文件名称, 文件内容
	file_list := strings.Split(string(file_data), "分割符")
	fmt.Println(len(file_list))
	if len(file_list) == 1 {
		fmt.Println("文件过大, 发送失败, 无法接收！！")
		return
	} else {
		//fmt.Println(file_list)
		file_name := file_list[0]
		file_save := file_list[1]
		// 创建文件对象
		file, _ := os.Create(`./` + file_name)
		// 向文件写入数据
		file.Write([]byte(file_save))
		fmt.Printf("文件%s接收成功", file_name)
	}
}
func main() {
	fmt.Println(`等待文件传输...`)
	// 监听发送端连接状况
	listen, err_lis := net.Listen(`tcp`, `127.0.0.1:8000`)
	if err_lis != nil {
		fmt.Println("err_lis =", err_lis)
		return
	}
	defer listen.Close()
	// 创建与发送端的连接对象
	conn, err_conn := listen.Accept()
	if err_conn != nil {
		fmt.Println("err_conn =", err_conn)
		return
	}
	defer conn.Close()
	//连接成功后, 接收发送端信息
	//定义变量接收发送文件信息
	file_data := make([]byte, 1024*102400)
	conn.Read(file_data)
	//fmt.Println(string(file_data))
	save_file(file_data)
	// 数据接收完毕, 向发送方发送提示信息
	send_list := `文件发送成功！！`
	conn.Write([]byte(send_list))
}
