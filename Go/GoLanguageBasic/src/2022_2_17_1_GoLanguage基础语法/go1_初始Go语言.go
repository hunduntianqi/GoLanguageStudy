package main

import (
	"fmt"
	"time"
)

/*
	什么是go语言:
		Go是一个开源的编程语言,可以很容易的构建简单, 可靠且高效的软件
	Go语言特点:
		1. 运行效率高,开发高效,部署简单
		2. 语言层面支持并发,易于利用多核实现并发
		3. 内置runtime(作用: 性能监控, GC等)
		4. 简单易学,丰富的标准库, 强大的网络库
		5. 内置强大的工具(gofmt), 跨平台编译, 内嵌C支持
	Go语言应用:
		1. 服务器编程:如处理日志,数据打包,虚拟机处理, 文件系统等
		2. 分布式系统: 数据库代理器,中间件等
		3. 网络编程: Web应用, API应用等
		4. 云平台
	Go语言命令行工具:
		1. go build: 用于编译源码文件, 代码包, 依赖包
		2. go run: 可以编译并运行Go源码文件
		3. go get: 可以用来动态获取远程代码包
		4. go install:
*/

func main() {
	fmt.Print("Hello World")
	time.Sleep(3 * time.Second)
}
