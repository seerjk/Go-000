package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func headers1(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		fmt.Fprintf(w, "%v: %v\n", name, headers)
		//for _, h := range headers {
		//	fmt.Fprintf(w, "%v: %v\n", name, h)
		//}
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/headers1", headers1)

	go http.ListenAndServe(":8090", nil)
	go http.ListenAndServe(":8070", nil)
	http.ListenAndServe(":8080", nil)
}
