package main

// 不使用struct，New时不取地址的Error

import (
	"errors"
	"fmt"
)

// Create a named type for our new error type.
// 没有使用struct
type errorString string

// Implement the error interface.
func (e errorString) Error() string {
	return string(e)
}

// New creates interface values of type error.
func New(text string) error {
	// 没有取地址
	return errorString(text)
}

// 自定义的Error对象
var ErrNamedType = New("EOF")

// 标准库的Error对象
var ErrStructType = errors.New("EOF")

func main() {
	if ErrNamedType == New("EOF") {
		// 自定义的Error，判断的string的值
		fmt.Println("Name Type Error")
	}

	if ErrStructType == errors.New("EOF") {
		// 标准库的Error，判断的&struct 地址
		// 文本内容一样时，地址也不会相等。避免2个不同error如果恰好字符串一样，被当作相等
		fmt.Println("Struct Type Error")
	}
}
