package main

import "fmt"

//组合接口
//发送接口
type Sender interface {
	Sende(string) error
}

//代码的复用
//网络连接 抽象（接口）
//open Send recive  Close

type Connection interface {
	Sender
	Open() error
	Recive() (string, error)
	Close() error
}
type TcpConnection struct {
}

func (c *TcpConnection) Open() error {

	return nil
}

func (c *TcpConnection) Recive() (string, error) {
	return "", nil
}

func (c *TcpConnection) Sende(msg string) error {
	return nil
}
func (c *TcpConnection) Close() error {
	return nil
}

//结构体组合接口

type Client struct {
	Connection
	C    Connection
	Name string
}

func main() {
	fmt.Println("11")
	var conn Connection = new(TcpConnection)
	fmt.Printf("%T, %#v\n", conn, conn)
	conn.Open()
	conn.Sende("xxxx")
	conn.Recive()
	conn.Close()
	client := Client{new(TcpConnection), new(TcpConnection), "KK"}
	client.Open()
	client.Sende("xxx")
	client.Recive()
	client.Close()

	client.C.Open()
	client.C.Sende("xxxx")
	client.C.Recive()
	client.C.Close()

}
