package main

import (
	"encoding/json"
	"fmt"
)

//json 是一个包，小写会获取不到值
/*
结构体属性的tag tagname:"tagvalue"
name,type,omitempty
name,omitempty
*/

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"-"`
	age      int    `json:"age"`
	Tel      string `json:"tel"`
	Addr     string `json:"addr"`
	IsBoy    bool   `json:"is_boy"`
}

//passworksjson不显示出来 `json _`
/*
Id => id => 结构体信息
Id =》 value
"id" => value
*/

func main() {
	user := User{Id: 1, Name: "kk", Password: "xxxx", age: 30, Tel: "123XXX", IsBoy: true}
	// encoding/json
	// 内存结构 => []byte (编码) 序列化
	// []byte, error = Marshal()
	// {"Id" : 1, "Name" : "kk", "Password" : "xxx", "Tel", "Add"}
	if ctx, err := json.Marshal(user); err == nil {
		fmt.Println(string(ctx))

	}
	//MarshalIndent 按格式实现
	if ctx, err := json.MarshalIndent(user, "", "\t"); err == nil {
		fmt.Println(string(ctx))
	}
	//[]byte =>内存结构（解码）反序列化
	ctx := `{"id":100,"Name":"kk2","Password":"xxx","Tel":"123xxx"}`
	var user2 User
	fmt.Println(json.Valid([]byte(ctx)))
	if err := json.Unmarshal([]byte(ctx), &user2); err == nil {
		fmt.Printf("%#v\n", user2)
	} else {
		fmt.Println(err)

	}

}
