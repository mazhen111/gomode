package main

/*

匿名组合体


 */
import (
	"fmt"
	"time"
)

type Addr struct {
	province string
	street string
	no string
}
type Tel struct {
	prefix string
	number string
	province string
}
//匿名组合方式
type User struct {
	id int
	name string
	*Addr   // 只指定类型 //会定义属性名Addr
	*Tel      // Tel
	birthday time.Time
	//province string
}

func main() {
	var user User
	user = User{Addr:&Addr{province: "马振"}}
	fmt.Printf("%T, %#v\n", user, user)
	fmt.Println(user.Addr.province)






}