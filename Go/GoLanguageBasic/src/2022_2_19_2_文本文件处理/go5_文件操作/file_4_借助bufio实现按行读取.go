/*
	bufio包实现了有缓冲的I/O:
		它包装一个io.Reader或io.Writer接口对象, 创建另一个也实现了该接口, 且同时还提供了缓冲和一些文本I/O的帮助函数的对象
		func NewReader(rd io.Reader) *Reader: 创建一个具有默认大小缓冲、从rd读取的*Reader
		func NewReaderSize(rd io.Reader, size int) *Reader: 创建一个具有最少有size尺寸的缓冲、从r读取的*Reader。如果
													参数r已经是一个具有足够大缓冲的* Reader类型值，会返回r
		func NewWriter(w io.Writer) *Writer: 创建一个具有默认大小缓冲、写入w的*Writer
		func NewWriterSize(w io.Writer, size int) *Writer: 创建一个具有最少有size尺寸的缓冲、写入w的*Writer。如果
													参数w已经是一个具有足够大缓冲的*Writer类型值，会返回w
		func (b *Reader) ReadBytes(delim byte) (line []byte, err error):读取直到第一次遇到delim字节，返回一个包含已
			读取的数据和delim字节的切片, 如果ReadBytes方法在读取到delim之前遇到了错误, 它会返回在错误之前读取的数据以及该错误（
			一般是io.EOF）, 当且仅当ReadBytes方法返回的切片不以delim结尾时，会返回一个非nil的错误
		func (b *Writer) WriteString(s string) (int, error): 写入一个字符串。返回写入的字节数,如果返回值nn < len(s),
															 还会返回一个错误说明原因
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	path := `D:\Users\Administrator\IdeaProjects\python_爬虫\2022_2_13_小说抓取\parse.txt`
	// 打开文件
	file, _ := os.Open(path)
	// 创建缓冲区
	buf := bufio.NewReader(file)
	// 按行读取数据
	for {
		data, _ := buf.ReadBytes('\n')
		fmt.Print(string(data))
		if string(data) == "" {
			break
		}

	}
}
