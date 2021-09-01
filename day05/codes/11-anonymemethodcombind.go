package main

import (
	"fmt"
	"time"
)

type Addr struct {
	province string
	street   string
	no       string
}
type Tel struct {
	prefix string
	number string
}

type User struct {
	id   int
	name string
	Addr
	Tel
	birthday time.Time
	province string
}

func (user User) SetProvince(name string) {
	user.name = name
	fmt.Println(user.name)

}

func main() {

	var user User
	user.name = "mazhen"
	user.SetProvince("hengs")
	fmt.Printf("%#v\n", user)

}
