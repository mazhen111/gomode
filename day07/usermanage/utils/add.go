package utils

import (
	"fmt"
	"time"
)

func Add() {
	fmt.Println("请输入要添加的用户信息")
	id := GetMaxId() + 1
	name := ""
	fmt.Println("请输入用户姓名")
	fmt.Scan(&name)
	addr := ""
	fmt.Printf("请输入%v的联系方式", name)
	fmt.Scan(&addr)
	tel := ""
	fmt.Printf("请输入%v的手机号码", name)
	fmt.Scan(&tel)
	birthday := ""
	fmt.Printf("请输入%v的生日，请按照中2006-01-02格式输入日期:", name)
	fmt.Scan(&birthday)
	_, err := time.Parse("2006-01-02", birthday)
	if err != nil {
		fmt.Println(err)
		fmt.Println("生日格式不正确无法解析")
		return

	}

}
