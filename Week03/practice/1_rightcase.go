package main

import (
	"fmt"
	"log"
	"net/http"
)

// 原则1：Keep yourself busy or do the work yourself

// 好的写法1：
func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello, GopherCon SG")
	})

	// badcase:
	//go func() {
	//	if err := http.ListenAndServe(":8080", nil); err != nil {
	//		log.Fatal(err)
	//		// log.Fatal() 本质执行的是 os.Exit(1)，导致 defer xxx 无法正常执行
	//	}
	//}()
	// goroutine(main())在从另一个goroutine(上面的👆)获取结果之前无法取得进展
	// 通常情况，main() 自己做这项工作（下面👇），会比委托它(go func(){}() )更简单
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
