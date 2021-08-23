package math

import "fmt"

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
