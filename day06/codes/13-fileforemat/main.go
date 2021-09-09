package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

// User /*
type User struct {
	ID   int
	Name string
}

func main() {
	fmt.Println("11")
	// json(文本), csv(文本), gob(二进制, go特有的，不能跨语言)
	user := []User{
		{1, "mazhen "},
		{2, "ma"},
	}
	//Register记录value下层具体值的类型和其名称。该名称将用来识别发送或接受接口类型值时下层的具体类型。
	//本函数只应在初始化时调用，如果类型和名字的映射不是一一对应的，会panic。
	gob.Register(User{}) //结构体的名称
	file, err := os.Create("users.gob")
	if err != nil {
		return

	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	//返回一个将编码后数据写入w的*Encoder
	encoder := gob.NewEncoder(file)
	//ncode方法将e编码后发送，并且会保证所有的类型信息都先发送
	fmt.Println(encoder.Encode(user))

}
