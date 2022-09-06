/*
	Socket:
		网络的Socket数据传输是一种特殊的I/O, Socket也是一种文件描述符, Socket也具有一个类似于打开文件的函数调用
		Socket():
			该函数返回一个整型的Socket描述符,随后的链接建立, 数据传输等操作都是通过该Socket实现的
		常用Socket类型:
			1. 流式(SOCK_STREAM)Socket:
				是一种面向链接的Socket, 针对于面向链接的TCP服务应用
			2. 数据报式(SOCK_DGRAM):
				是一种无连接的Socket, 对应于无连接的UDP服务应用
*/

package main

func main() {

}
