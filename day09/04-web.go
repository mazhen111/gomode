package main

import (
	"fmt"
	"net/http"
)

// 处理器函数
func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi"))

}

type Hellp struct {
}

func (h *Hellp) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	fmt.Fprint(writer, "help")
}

func main() {
	// url => 找处理器函数 => 调用处理器函数(http包)
	// 指定url和处理器关系
	// 处理器函数签名由http包定义
	add := ":8888"
	http.HandleFunc("/home", Home)
	http.Handle("/help", new(Hellp))
	http.ListenAndServe(add, nil)

}
