package main

import (
	"fmt"
)

//匿名接口断言

type User struct {
	Name     string
	Password string
}

func main() {
	fmt.Println("11")
	var empty interface{}
	empty = User{Name: "x", Password: "y"}
	if u, ok := empty.(User); ok {

		fmt.Println(u.Name, u.Password)
	}
	fmt.Printf("%T, %#v\n", empty, empty)

}
