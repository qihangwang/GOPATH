package main

import (
	"fmt"
	"net/http"
)

//创建处理器函数
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!", r.URL.Path)
}

func main() {
	//适配器
	http.HandleFunc("/http", handler)
	//HandleFunc注册一个处理器handler和对应的模式pattern
	//创建路由
	http.ListenAndServe(":8080", nil)
	// listenAndServer监听TCP地址addr 并且会使用handler参数调用Server函数处理接收到的连接。handler参数一般会设为nil，此时会使用DefaultServeMux（多路复用器）
}
