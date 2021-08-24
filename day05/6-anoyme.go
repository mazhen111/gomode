package main

import (
	"fmt"
	"time"
)

//匿名结构体
func main() {
	//定义匿名结构体的变量
	var user struct{
		id       int
		name     string
		tel      string
		addr     string
		birthday time.Time
	}
	fmt.Printf("%T, %#v\n", user, user)
	user.id = 10
	user.name = "kk"
	fmt.Printf("%T, %#v\n", user, user)

	// 初始化
	// 零值
	// 字面量
	user = struct {
		id       int
		name     string
		tel      string
		addr     string
		birthday time.Time
	}{10, "kk", "xxx", "xxxxx", time.Now()} //结构体类型{}

	fmt.Printf("%T, %#v\n", user, user)

}
