package utils

import "fmt"

func SetPasswd(name string) string {
	//定义设置密码函数
	for {
		//对两次输入的密码进行校验，两次一致时再进行md5转换存储
		fmt.Printf("请输入%v账号新密码", name)

	}

}
