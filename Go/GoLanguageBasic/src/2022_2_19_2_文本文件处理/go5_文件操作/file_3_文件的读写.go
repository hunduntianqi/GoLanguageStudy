/*
	str := fmt.Sprintf(格式化字符串): 将格式化后的字符串赋值给str
*/

package main

import (
	"fmt"
	"os"
)

// 定义函数向文件写入数据
func WriteFile(path string) {
	// 打开文件
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	// 向文件写入数据
	str := fmt.Sprintf("%s", "this is one txt file")
	file.WriteString(str)
	os.Stdout.WriteString("this is one txt file\n")
	// 关闭文件
	file.Close()
}

// 定义函数读取文件数据
func ReadFile(path string) {
	// 打开文件
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("err=", err)
	}
	print(file.Name() + "\n")
	// 定义字节切片
	file_data := make([]byte, 1024, 1024)
	data, err_read := file.Read(file_data) // 返回文件数据的字节数
	if err_read != nil {
		fmt.Println("err_read=", err_read)
		return
	}
	fmt.Println("文件字节数=", data)
	fmt.Println("文件内容为:", string(file_data))
	file.Close()
}

// 定义函数删除文件
func RemoveFile(path string) {
	os.Remove(path)
}
func main() {
	// 定义文件路径
	path := "./test.txt"
	// 调用函数写入数据
	WriteFile(path)
	// 调用函数读取文件
	ReadFile(path)
	// 调用函数删除文件
	RemoveFile(path)

}
