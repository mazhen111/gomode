package main

import (
	"fmt"
	"net"
	"time"
)

//并发网络编程
func main() {
	addr := "127.0.0.1:8888"
	protocol := "tcp"
	listener, err := net.Listen(protocol, addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(conn)
			continue
		}

		go func() {
			fmt.Println("客户端链接成功:", conn.RemoteAddr())
			time.Sleep(10 * time.Second)
			//发送时间
			fmt.Fprintln(conn, time.Now().Format("2006-01-02 15:04:05"))
			conn.Close()
			fmt.Println("客户端退出", conn.RemoteAddr())

		}()
	}

}
