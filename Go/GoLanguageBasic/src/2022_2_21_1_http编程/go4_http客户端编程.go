package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://pic.netbian.com/4kmeinv/index_3.html")
	if err != nil {
		fmt.Println("err =", err)
		return
	}
	defer resp.Body.Close()
	// 打印请求状态码
	fmt.Println("status =", resp.StatusCode)
	fmt.Println("status =", resp.Status)
	fmt.Println("header =", resp.Header)
	text := make([]byte, 1024*4)
	var tmp string
	for {
		n, err := resp.Body.Read(text)
		if n == 0 {
			fmt.Println("read err =", err)
			break
		}
		tmp += string(text)
	}
	fmt.Println()

}
