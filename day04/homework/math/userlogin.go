package math

import (
	"crypto/md5"
	"fmt"
)

var  (
	username string
	password string
	userinfo=[]map[string]string{}
)
func Userlogin()  {
	var   logbool bool
	//用户登录验证
	userinfos:=map[string]string{
		"username":"",
		"password":"",

	}
	if len(userinfo) == 0{
		fmt.Println("注册用户")
		fmt.Println("请输入用户")
		fmt.Scanln(&username)
		fmt.Println("请输入密码")
		fmt.Scanln(&password)
		password = fmt.Sprintf("%X",md5.Sum([]byte(password)))
		fmt.Println(password)
		userinfos["username"]=username
		userinfos["password"]=password
		userinfo=append(userinfo,userinfos)
		fmt.Println(userinfo)
	}
	fmt.Println("登录用户")
	for i := 0; i < 3; i++ {
		fmt.Println("请输入用户")
		fmt.Scanln(&username)
		fmt.Println("请输入密码")
		fmt.Scanln(&password)
		password=fmt.Sprintf("%X", md5.Sum([]byte(password)))
		fmt.Println(password)
		for _, v := range userinfo {
			fmt.Println(v["username"])
			fmt.Println(v["password"])
			if v["username"] == username && v["password"] ==password{
				fmt.Println("登录成功")
				logbool=true

			} else {
				fmt.Println("用户名或密码错误")

			}

		}
		if logbool {
			break
		}else if  i == 2  {
			fmt.Println("重试次数达到上限")
		}

	}

}