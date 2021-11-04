package main

import "fmt"

//空接口
type Empty interface {
}

type User struct {
	Name     string
	Password string
}

func main() {
	var empty Empty
	empty = 1
	fmt.Println(empty)
	if v, ok := empty.(int); ok {
		fmt.Println(v)
	}
	empty = "xxx"
	fmt.Println(empty)
	empty = true
	fmt.Printf("%T, %#v\n", empty, empty)
	empty = User{Name: "mazhen", Password: "1111"}
	if v, ok := empty.(User); ok {
		fmt.Println(v.Name, v.Password)

	}
	// empty.Name
	fmt.Printf("%T, %#v\n", empty, empty)

}
