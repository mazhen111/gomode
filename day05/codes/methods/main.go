package main

import (
	"fmt"
	"methods/models"
)

func main() {
	fmt.Println("11")
	user := models.NewUser(1, 11, "22")
	fmt.Printf("%#v\n", user)
	models.AddAge(user) // 修改属性
	fmt.Printf("%#v\n", user)
	fmt.Println(models.GrtName(user))
	puser := &user // 取引用 取地址
	// 调用函数如何调用
	(*puser).AddAge() // 解引用 取值
	fmt.Println(*puser)
	fmt.Println((*puser).GetName())
}
