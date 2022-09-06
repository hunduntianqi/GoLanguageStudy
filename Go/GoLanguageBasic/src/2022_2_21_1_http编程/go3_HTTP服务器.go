/*
	HTTP编程:
		Go语言标准库内建提供了net / http包, 涵盖了客户端和服务端的具体实现
*/

package main

import (
	"fmt"
	"net/http"
)

// w: 给客户端回复数据, req: 读取客户端发送的数据
func HandConn(w http.ResponseWriter, req *http.Request) {
	// 给客户端回复数据
	w.Write([]byte(`hello world`))
	// 获取用户请求网址
	fmt.Println("req.url =", req.URL)
	// 获取用户请求的方法类型
	fmt.Println("req.method =", req.Method)
	// 获取用户请求头信息
	fmt.Println("req.header =", req.Header)
}

func main() {
	// 注册处理函数, 用户连接, 自动调用指定的处理函数
	http.HandleFunc("/", HandConn)
	// 监听绑定
	http.ListenAndServe(":8000", nil)
}
