package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// 定义函数复制文件
func Copy(file_path, result_path string, ch *chan string) {
	//fmt.Println(strings.Split(file_path, `\`))
	// 只读方式打开文件对象
	file, _ := os.Open(file_path)
	// 获取文件信息
	f_data, _ := os.Stat(file_path)
	// 定义与文件大小一致的字节切片
	data_byte := make([]byte, f_data.Size(), f_data.Size()+100)
	file.Read(data_byte)
	//fmt.Println(string(data_byte))
	// 拼接目标文件路径
	result_path_file := result_path + `\\` + strings.Split(file_path, `\`)[len(strings.Split(file_path, `\`))-1]
	// 创建目标文件对象
	result_file, _ := os.Create(result_path_file)
	// 创建写入缓冲数据
	// 写入数据
	result_file.Write(data_byte)
	//time.Sleep(1 * time.Second)
	fmt.Printf("图片%s已复制成功\n", f_data.Name())
	*ch <- file_path

}

func main() {
	n := runtime.GOMAXPROCS(2)
	fmt.Println(n)
	start := time.Now()
	// 创建文件夹保存复制的文件
	var result_path string
	fmt.Println("请输入目标文件夹路径")
	fmt.Scan(&result_path)
	os.Mkdir(result_path, 0666)
	// 定义源文件夹路径变量
	var path string
	fmt.Println("请输入源文件路径:")
	fmt.Scan(&path)
	path = path + "\\*" // "*"代表获取文件夹下所有文件路径
	// 调用函数获取该文件夹下所有文件路径
	files, _ := filepath.Glob(path)
	ch := make(chan string)
	for _, file_path := range files {
		go Copy(file_path, result_path, &ch)
	}
	for i := 0; i < len(files); i++ {
		if i == len(files)-1 {
			break
		} else {
			<-ch
		}

	}
	end := time.Now()
	fmt.Println(end.Sub(start))
}
