package math

import (
	"fmt"
	"strconv"
)

var (
id    string
name  string
phone string
addr  string
users = []map[string]string{}
)
func AddUser()  {
	userinfo := map[string]string{
		"id":    "",
		"name":  "",
		"phone": "",
		"addr":  "",
	}
	fmt.Println(userinfo)
	fmt.Println("请输入姓名")
	fmt.Scanln(&name)
	fmt.Println("请输入手机号")
	fmt.Scanln(&phone)
	fmt.Println("请输入通讯地址")
	fmt.Scanln(&addr)
	fmt.Println(len(users))
	if len(users) == 0{
		id = "0"

	}else{

		id = strconv.Itoa(len(users))
		fmt.Println(id)
	}
	userinfo["id"]=id
	userinfo["name"]=name
	userinfo["phone"]=phone
	userinfo["addr"]=addr
	users=append(users,userinfo)
	fmt.Println(users)
}