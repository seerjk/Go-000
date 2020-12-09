package main

import (
	"fmt"
	"log"
	"net/http"
)

// 原则1：Keep yourself busy or do the work yourself

// 不好的写法1：
func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello, GopherCon SG")
	})
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal(err)
			// log.Fatal() 本质执行的是 os.Exit(1)，导致 defer xxx 无法正常执行
		}
	}()

	select {} // 空的select{}语句将永远阻塞
	// 不建议这样的写法，原因是main()无法感知goroutine什么时候退出
}
