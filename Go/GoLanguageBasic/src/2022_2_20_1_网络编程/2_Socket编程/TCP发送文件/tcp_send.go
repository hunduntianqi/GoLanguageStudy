package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// 创建tcp客户端连接对象
	conn, err_conn := net.Dial(`tcp`, `127.0.0.1:8000`)
	if err_conn != nil {
		fmt.Println("err_conn=", err_conn)
		return
	}
	// 定义变量存储文件路径
	var file_path string
	fmt.Println("请输入要发送的文件路径:")
	fmt.Scan(&file_path)
	// 创建文件对象, 打开文件
	file, err := os.Open(file_path)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	// 调用os.Stat函数, 获取文件信息
	f_type, err_type := os.Stat(file_path)
	if err_type != nil {
		fmt.Println("err_type=", err_type)
		return
	}
	if f_type.Size() > 1024*102400 {
		fmt.Printf("文件大小为%d", f_type.Size())
		fmt.Println("不允许发送超过100M大小的文件！！")
		return
	}
	// 定义切片存储文件信息
	save_data := make([]byte, f_type.Size())
	// 定义字符串存储文件名称
	file_name := f_type.Name()
	fmt.Println(file_name)
	// 读取文件信息到切片中
	file.Read(save_data)
	// 拼接文件信息
	file_data := file_name + `分割符` + string(save_data)
	//向接收段发送数据
	//发送文件内容
	conn.Write([]byte(file_data))
	// 接收文件接收方的回复, 确保文件接收成功
	receive := make([]byte, 1024)
	conn.Read(receive)
	fmt.Println("receive =", string(receive))
	// 数据发送完毕, 关闭通道
	conn.Close()
}
