/*
	error接口:
		Go语言引入了一个关于错误处理的标准模式, 即error接口, 是Go语言内建的接口类型
		该接口定义如下:
			type error interface{
				Error() string
			}
		Go语言标准库代码包errors为用户提供如下方法:
			package errors
			type errorString struct {
				text string
			}

			func New(text string) error {
				return &errorString(text)
			}

			func (e *errorString) Error() string {
				return e.text
			}

		另一个可以生成error类型值的方法是调用fmt包中的Errorf函数:
			package fmt
			import "errors"

			func Errorf(format string, args ...interface{}) error {
				return errors.New(Sprintf(format, args...))
			}
		error接口应用:

*/

package main

import (
	"errors"
	"fmt"
)

func Div(a, b int) (result int, err error) {
	if b == 0 {
		err = errors.New("分母不能为0")
	} else {
		result = a / b
	}
	return
}

func main() {
	//var err error = fmt.Errorf("%s", "this is normal error")
	// fmt.Errof()生成错误类型变量
	err1 := fmt.Errorf("%s", "this is normal error1")
	fmt.Println(err1)
	// errors.New()生成错误类型变量
	err2 := errors.New("this is normal error2")
	fmt.Println(err2)

	result, err := Div(10, 0)
	if err != nil {
		fmt.Printf("异常:%s", err)
	} else {
		fmt.Printf("result=%d\n", result)
	}

}
