package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

/*
	爬取步骤:
		1. 明确目标(要知道在哪个网站或范围进行搜索)
		2. 爬: 将网站的内容爬取下来
		3. 取: 对数据进行清洗筛选
		4. 处理数据(持久化存储)
*/
func SpiderBaiDu(i int, page chan<- int) {
	fmt.Println("开始爬取第" + strconv.FormatInt(int64(i), 10) + "页数据...")
	if i == 1 {
		url := "https://pic.netbian.com/4kmeinv/"
		// 向服务器发送请求, 获取响应对象
		resop, err := http.Get(url)
		if err != nil {
			fmt.Println("err =", err)
			return
		}
		fmt.Println("Status =", resop.Status)
		// 打印响应状态码
		fmt.Println("StatusCode =", resop.StatusCode)
		// 创建文件, 保存页面源码数据
		file, err_file := os.Create("src/2022_2_21_1_http编程/百度贴吧/" + strconv.FormatInt(int64(i), 10) + "_百度贴吧.html")
		if err_file != nil {
			fmt.Print("err_file =", err_file)
			return
		}
		// 打印页面源码
		html := make([]byte, 1024*8)
		var text string
		for {
			n, err := resop.Body.Read(html)
			if err != nil {
				fmt.Println("err =", err)
				break
			} else {
				file.Write(html)
				text += string(html[:n])
				//fmt.Println(text)
			}
		}
		file.Close()
		fmt.Println("第" + strconv.FormatInt(int64(i), 10) + "页数据抓取完毕")
		page <- i
	} else {
		url := "https://pic.netbian.com/4kmeinv/index_" + strconv.FormatInt(int64((i-1)*50), 10) + ".html"
		// 向服务器发送请求, 获取响应对象
		resop, err := http.Get(url)
		if err != nil {
			fmt.Println("err =", err)
			return
		}
		fmt.Println("Status =", resop.Status)
		// 打印响应状态码
		fmt.Println("StatusCode =", resop.StatusCode)
		// 创建文件, 保存页面源码数据
		file, err_file := os.Create("src/2022_2_21_1_http编程/百度贴吧/" + strconv.FormatInt(int64(i), 10) + "_百度贴吧.html")
		if err_file != nil {
			fmt.Print("err_file =", err_file)
			return
		}
		// 打印页面源码
		html := make([]byte, 1024*8)
		var text string
		for {
			n, err := resop.Body.Read(html)
			if err != nil {
				fmt.Println("err =", err)
				break
			} else {
				file.Write(html)
				text += string(html[:n])
				//fmt.Println(text)
			}
		}
		file.Close()
		fmt.Println("第" + strconv.FormatInt(int64(i), 10) + "页数据抓取完毕")
		page <- i
	}

}

func DoWork(start, end int) {
	var page chan int
	for i := start; i <= end; i++ {
		go SpiderBaiDu(i, page)
	}

	for i := start; i <= end; i++ {
		//if _, ok := <-page; ok {
		//	//fmt.Println("ok =", ok)
		//	fmt.Printf("ok =%d, 第%d页数据爬取完毕", ok, <-page)
		//}
		//fmt.Println("第" + string(<-page) + "页数据抓取完毕")
		<-page
	}

}

func main() {
	var start, end int
	fmt.Println("请输入爬取的起始页(>=1):")
	fmt.Scan(&start)
	fmt.Println("请输入爬取的终止页:")
	fmt.Scan(&end)
	os.Mkdir("src/2022_2_21_1_http编程/百度贴吧", 0666)
	// 定义函数执行抓取功能
	DoWork(start, end)
}
