package main

import (
	"fmt"
	"time"
)
//组合---》继承

type Addr struct {
	province string
	street string
	no string
}

type Tel struct {
	prefix string
	mumber string
}
//命名组合
type User struct {
	id int
	name string
	addr Addr
	tel  Tel
	pirthday time.Time
}

func main() {
	var user User
	user=User{addr: Addr{province:"陕西"}}
	fmt.Printf("%T,%#v\n",user,user)
	fmt.Println(user.addr.province)

	//继承结构体重新赋值
	user.addr.province="北京"
	fmt.Println(user.addr.province)
}


