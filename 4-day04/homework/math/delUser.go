package math

import (
	"fmt"
	"strconv"
)

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