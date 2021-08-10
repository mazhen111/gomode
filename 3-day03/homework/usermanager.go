package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var info string
var (
	id    string
	name  string
	phone string
	addr  string
	users = []map[string]string{}


)
func addUser()  {
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

func delUser ()  {
	fmt.Println("delUser")
	fmt.Println(users)
	if len(users) ==0{
		fmt.Println("用户为空")
		return

	}
	var deleid string
	fmt.Println("请选择删除用户的id")
	fmt.Scanln(&deleid)
	for i,v :=range users{
		if v["id"] == deleid{
			fmt.Println(v)
			var ConfirmDeletion string
			fmt.Println("输入y/n确认删除")
			fmt.Scanln(&ConfirmDeletion)
			if ConfirmDeletion == "y" || ConfirmDeletion =="y"{
				if deleid == "0"{
					users = users[1:]

				}else if deleid == strconv.Itoa(len(users)){
					users =users[:len(users)-1]

				}else {

					users = append(users[:i :i+1 ])
				}

			}else {
				fmt.Println("取消删除")
				return
			}
			fmt.Println(users)




		}else {
			fmt.Println("输入的ID不正确")
		}
		//fmt.Println(i,v)
		//fmt.Println(v["id"])
	}




}

func modifyUser()  {
	if len(users) == 0{
		fmt.Println("用户不存在")

	}
	var modid string
	fmt.Println("请输入修改的ID")
	fmt.Scanln(&modid)
	for _ ,v :=range users{
		if v["id"] ==id{
			fmt.Printf("修改的用户信息为%q\n",v)
			fmt.Println("输入姓名")
			fmt.Scanln(&name)
			fmt.Println("phone")
			fmt.Scanln(&phone)
			fmt.Println("输入地址")
			fmt.Scanln(&addr)
			v["name"]=name
			v["phone"]=phone
			v["addr"]=addr
			fmt.Println(addr)

		}else {
			fmt.Println("输入的ID不正确")
		}

	}
	fmt.Println("modifyUser")
	fmt.Println(users)
}

func queryUser()  {
	fmt.Println("queryUser")
	if len(users) == 0{
		fmt.Println("用户为空")
		return
	}
	var querid string
	fmt.Println("请输入修改的ID")
	fmt.Scanln(&querid)
	for _,v :=range users{
		if v["id"] ==id{
			fmt.Println(v)
		}else {
			fmt.Println("输入的ID不正确")


		}





	}



}


func main() {


	info =`
    ---用户管理---
1、添加用户
2、删除用户
3、修改用户
4、查找用户
5、退出(q键)
`

	for true{
		//fmt.Printf("%v", users)
		//fmt.Println(id,name,phone,addr,users)
		fmt.Println(info)
		var option string
		fmt.Print("请选择：")
		fmt.Scanln(&option)
		if option == "1"{
			addUser()
			time.Sleep(300)

		}else if option =="2" {
			delUser()
		}else if option =="3" {
			modifyUser()
		}else if option == "4"{
			queryUser()
		}else if option == "q" || option == "Q"{
			os.Exit(0)
		}else {
			fmt.Println("请选择(1/2/3/4)")
		}


	}



	}








