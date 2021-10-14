package main

import (
	"fmt"
	"log"
	"net/http"
)

// Engine 类实现 Handler 接口，用来处理所有 http 请求
type Engine struct {}


// 实现 Handler 接口中的 ServeHTTP 函数
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request)  {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

func main() {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":9999", engine))
}