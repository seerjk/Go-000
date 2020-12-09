package main

import "fmt"

// 和标准库一样使用struct
type errorString struct {
	s string
}

func (e errorString) Error() string {
	return e.s
}

func NewError(text string) error {
	// 没有取 &struct 地址
	return errorString{text}
}

var ErrType = NewError("EOF")

func main() {
	if ErrType == NewError("EOF") {
		// 判断相等时，看struct中每个字段是否相等
		// 只有一个string字段，值都是EOF
		// 所以，必须像标准库一样取地址，判断内存地址指向是否相等
		fmt.Println("Error:", ErrType)
	}
}
