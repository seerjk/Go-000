package main

import "errors"

func rpc1() (int, error) {
	return 1, errors.New("xxx1")
}

func rpc2() (int, error) {
	return 2, errors.New("xxx2")
}

type Result struct {
	Value int
	Err   error
}

func main() {
	// 微服务并发 call 多个下游rpc 都方式goroutine时
	var a, b int
	var err1, err2 error

	var ch chan Result

	go func() {
		// call RPC1
		// 产生的error无法return回来
		a, err1 = rpc1()
		ch <- Result{a, err1}
	}()

	go func() {
		// call RPC2
		// 产生的error无法return回来
		b, err2 = rpc2()
	}()

	// 或者用chan把 error传回来
	// 不够好用
	<-ch
}
