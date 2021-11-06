package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Id       int
	Name     string
	Password string
	Tel      int
	Addr     string
	IsBoy    bool
}

func main() {
	//user :=[]User{
	//	{1, "kk", "xxx", 30,  "xxx", true},
	//	{2, "kk2", "xxx", 30,  "xxx", true},
	//	{3, "kk3", "xxx", 30,  "xxx", true},
	//}
	//
	//fmt.Printf("11")
	//file,_:=os.Create("user.json")
	//defer file.Close()
	//encoder:=json.NewEncoder(file)
	//encoder.Encode(user)
	file, err := os.Open("user.json")
	fmt.Println(err)
	defer file.Close()
	decode := json.NewDecoder(file)
	var user2 []User
	decode.Decode(&user2)
	fmt.Println(user2)

}
