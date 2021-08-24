package main

import (
	"fmt"
	"time"
)

/*
指针结构体

 */

type User struct {
	id int
	name string
	addr string
	tel  string
	birthday time.Time
}
//new函数
func NewUser(id int,name,addr,tel string,birthday time.Time) *User {
	return &User{id,name,addr,tel,birthday}
}
func main() {
	var user =&User{id: 10, name: "kk"}
	fmt.Printf("%T,%#v\n",user,user)
	//new函数
	user = new(User)
	fmt.Printf("%T,%#v\n",user,user)
	user = NewUser(1, "kk", "", "", time.Now())
	fmt.Printf("%T,%#v\n",user,user)
	//属性访问
	fmt.Printf(user.name)
	//修改
	user.name="mazhen"
	fmt.Printf("%#v\n",user)




}
