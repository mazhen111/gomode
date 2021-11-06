package main

import "fmt"

type User struct {
	Id       int
	Nane     string
	Password string
}

func (user User) String() string {
	return fmt.Sprintf("User[Name=%s]", user.Nane)

}

func main() {
	//结构体格式化输出
	fmt.Println("11")
	var user User = User{0, "xxx", "yyy"}
	fmt.Println(user)
	puser := &User{Id: 1, Nane: "xxxx", Password: "yyy"}
	fmt.Println(puser)
}
