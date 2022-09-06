/*
	JSON:
		是一种比XML更轻量级的数据交换格式, 再易于人们阅读和编写的同时, 也易于程序解析和生成,
	JSOn采用完全独立于编程语言的文本格式, 表现为键值对集合的文本描述形式, 使JSON成为较为理想的
	, 跨平台, 跨语言的数据交换语言
		开发者可以用JSON传输简单的字符串, 数字, 布尔值, 也可以传输一个数组, 或者一个更复杂的符合结构,
	再web开发领域中, JSON被广泛应用于Web服务端和客户端之间的数据通信
    Go语言对Json的支持:
		使用Go语言内置的encoding / json标准库, 开发者可以轻松使用Go程序生成和解析JSON格式的数据
	编码json:
		1. 通过结构体生成JSON
			注意:结构体中字段名首字母必须大写, 否则无法读取字段！！
			使用json.Marshal()函数可以对一组数据进行JSON格式编码, json.Marshal()函数声明如下:
				func Marshal(v interface{}) ([]byte, error)
			格式化输出json:
				func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)
			struct_tag使用:
				结构体二次编码:
					在结构体字段后加入语句:`json:"字段名"`, 则输出字段会按照此字段名输出, 此操作属于二次编码
				指定不输出字段:
					在结构体字段后加入语句:`json:"-"`
				修改指定字段数据类型:
					在结构体字段后加入语句:`json:",数据类型"`
		2. 通过map生成JSON:
			2.1 先定义一个map数据, 并初始化
			2.2 调用json.Marshal()函数或json.MarshalIndent()函数生成json数据
	解码JSON:
		func Unmarshal (data []byte, v interface{}) error {}
		1. 解码到结构体:
			func Unmarshal(daya []byte, &结构体变量) error {}
			注意: 传入结构体变量必须为取地址, 否则无法修改实际结构体变量内容
		2. 解码到map
			func Unmarshal(daya []byte, &map数据变量) error {}
			注意: 传入map数据变量必须为取地址, 否则无法修改实际结构体变量内容
*/

package main

import (
	"encoding/json"
	"fmt"
)

// 定义Json测试结构体
type json_test struct {
	Company string `json:"公司"`
	Addr    string `json:"地址"`
	Persons int    `json:",string"` // 将原来为整型的字段Persons转换为字符串输出
}

func main() {
	/*
		通过结构体生成json
	*/
	// 结构体变量初始化
	var js1 json_test = json_test{Company: "富士康", Addr: "河南省郑州市航空港区综合保税区", Persons: 100000}
	// 调用json.Marshal函数生成json格式数据
	js_result1, error := json.Marshal(js1)
	if error == nil {
		fmt.Println(string(js_result1))
	}

	// 调用json.MarshalIndent函数格式化输出json
	js_result2, error2 := json.MarshalIndent(js1, ``, ` `)
	if error2 == nil {
		fmt.Println(string(js_result2))
	}
	// 解析json到结构体
	// 定义结构体变量
	var js_decode json_test
	err_decode := json.Unmarshal(js_result1, &js_decode)
	if err_decode == nil {
		fmt.Println(js_decode)
	} else {
		fmt.Println(err_decode)
	}

	/*
		通过map生成json
	*/
	// 定义一个map
	json_map := make(map[string]interface{})
	json_map["姓名"] = `郭鹏涛`
	json_map["性别"] = `男`
	json_map["公司"] = "富士康"
	json_map["地址"] = "河南省郑州市航空港区综合保税区"
	json_map["人数"] = "10000-300000"
	json_map["所属部门"] = "idpbg事业群-MLB事业处-测试工程处-失效分析工程部-维修改善三课-RF组"
	// 调用json.Marshal()函数生成json数据
	js_result3, err3 := json.Marshal(json_map)
	if err3 == nil {
		fmt.Println(string(js_result3))
	}
	// 格式化输出
	js_result4, err4 := json.MarshalIndent(json_map, ``, ` `)
	if err4 == nil {
		fmt.Println(string(js_result4))
	}
	// 解析json到map
	//定义map数据变量
	var map_decode map[string]interface{}
	err_map_decode := json.Unmarshal(js_result3, &map_decode)
	if err_map_decode == nil {
		fmt.Println("人数=", map_decode["人数"])
	}
}
