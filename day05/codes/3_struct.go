package main

import (
	"fmt"
	"time"
)

/*
结构体定义使用 struct 标识， 需要指定其包含的属性（ 名和类型），在定义结构体时可以为
结构体指定结构体名（命名结构体），用于后续声明结构体变量使用
*/
//id =int
//name =string
//addr = string
// tel => string
// birthday => time.Time
type User struct {
	id       int
	name     string
	addr     string
	tel      string
	birthday time.Time
}

func main() {
	var user User
	// 零值是由个元素的零值组成的一个结构体的变量
	fmt.Printf("%T,%#v\n", user, user)
	// 字面量(按照属性顺序定义的字面量)
	// 必须严格按照结构体属性定义顺序指定
	// 每个属性都必须指定对应的值
	//user=User{10, "kk", "西安", "xxxxx", time.Now()}
	fmt.Printf("%#v\n", user)
	// 按照属性名定义的字面量
	user = User{id: 1, name: "马振"}
	fmt.Printf("%#v\n", user)
	// 属性访问和修改
	fmt.Println(user.id)
	user.id = 100000
	fmt.Println(user.id)
	fmt.Printf("%#v", user)

}
