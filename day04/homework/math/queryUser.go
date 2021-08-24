package math

import "fmt"

func QueryUser()  {
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
