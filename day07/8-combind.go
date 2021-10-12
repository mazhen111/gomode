package main

import "fmt"

//组合接口
//发送接口
type Sender interface {
	Sender(string) error
}

//代码的复用
//网络连接 抽象（接口）
//open Send recive  Close

type Connection interface {
	Sender
	Open()
	Recive() (string error)
	Close() error
}

func main() {
	fmt.Println("11")

}
