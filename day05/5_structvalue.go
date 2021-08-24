package main

import (
	"fmt"
	"time"
)
// 指针，地址，切片，映射
type User struct {
	id int
	name string
	addr string
	tel  string
	birthday time.Time
}

//结构体复制会申请新的内存，不会影响别的变量
func main() {
	var user User
	fmt.Printf("%#v",user)
	c1 := user
	c1.name="js"
	fmt.Printf("%#v\n",user)
	fmt.Printf("%#v\n",c1)


}

