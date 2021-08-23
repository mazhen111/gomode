package main

import (
	"fmt"
	"homework/math"
	"os"
	"time"
)
func main() {
	info := `
 ---用户管理---
1、添加用户
2、删除用户
3、修改用户
4、查找用户
5、退出(q键)
`
math.Userlogin()
	for true{
		//fmt.Printf("%v", users)
		//fmt.Println(id,name,phone,addr,users)
		fmt.Println(info)
		var option string
		fmt.Print("请选择：")
		fmt.Scanln(&option)
		if option == "1"{
			math.AddUser()
			time.Sleep(300)

		}else if option =="2" {
			math.DelUser()
		}else if option =="3" {
			math.ModifyUser()
		}else if option == "4"{
			math.QueryUser()
		}else if option == "q" || option == "Q"{
			os.Exit(0)
		}else {
			fmt.Println("请选择(1/2/3/4)")
		}


	}

}
