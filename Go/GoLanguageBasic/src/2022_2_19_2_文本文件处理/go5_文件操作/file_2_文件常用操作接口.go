/*
	创建文件对象:
		1. func Create(name string) (file *File, err error): 根据文件名name创建文件, 返回可读写的文件对象, 如果文件已存在会
															 清空文件数据, 默认权限为0666
		2. func NewFile(fd uintptr, name string) *File:根据文件描述符创建相应的文件, 返回一个文件对象
	打开文件:
		1. func Open(name string) (file *File, err Error): 以只读方式打开一个名称为name的文件, 内部实现其实调用了OpenFile
		2. func OpenFile(name string, flag int, perm uint32) (file *File, err Erroe):打开名称为name的文件, flag为打开方式
																					  包括只读, 读写等, perm是权限
			flag参数类型:
				O_RDONLY	只读方式打开
				O_WRONLY	只写方式打开
				O_RDWR	   读写方式打开
				O_APPEND	追加方式打开
				O_CREATE	不存在，则创建
				O_EXCL	如果文件存在，且标定了O_CREATE的话，则产生一个错误
				O_TRUNG	如果文件存在，且它成功地被打开为只写或读写方式，将其长度裁剪唯一。（覆盖）
				O_NOCTTY	如果文件名代表一个终端设备，则不把该设备设为调用进程的控制设备
				O_NONBLOCK	如果文件名代表一个FIFO,或一个块设备，字符设备文件，则在以后的文件及I/O操作中置为非阻塞模式。
				O_SYNC	当进行一系列写操作时，每次都要等待上次的I/O操作完成再进行
	写文件:
		1. func (file *File) Write(b []byte) (n int, err Error):写入byte类型的信息到文件
		2. func (file *File) WriteAt(b []byte, off int64) (n int, err Error):在指定位置开始写入byte信息
		3. func (file *File) WeiteString(s string) (ret int, err Erroe): 写入string信息到文件
	读文件:
		1. func (file *File) Read(b []byte) (n int, err Error): 读取数据到字节数组b中
		2. func (file *File) ReadAt(b []byte, off int64) (n int, err Error):从off开始读取数据到字节数组b中
	删除文件:
		1. func Remove(name string) error:删除name指定的文件或文件夹
		2. func RemoveAll(path string) error: 删除path指定的文件，或目录及它包含的任何下级对象, 如果path指定的对象不存在,
											  RemoveAll会返回nil而不返回错误
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	file_data := make([]byte, 1024, 1024)
	// 打开文件
	file, err_file := os.Open("D:\\Users\\Administrator\\IdeaProjects\\python_爬虫\\2022_2_13_小说抓取\\parse.txt")
	if err_file == nil {
		file_size, _ := file.Read(file_data)
		os.Stdout.WriteString("这是读取文件的内容:")
		fmt.Println("文件字节数为:", file_size)
		fmt.Println(string(file_data))
	}

}
