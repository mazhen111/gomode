package main

import "fmt"

// 见类型知含义 & 可以添加方法
// 现有数据类型
type Counter int
type User map[string]string
type Callback func() error

func main() {
	var counter Counter
	fmt.Printf("%T,%v\n", counter, counter)
	counter += 10
	fmt.Printf("%T,%v\n", counter, counter)
	var num int = 10
	c2 := counter + Counter(num)
	fmt.Printf("%T,%v\n", c2, c2)
	var user User = make(User)
	fmt.Printf("%T,%T\n", user, user)
	user["id"] = "11"
	fmt.Println(user)
	callbacks := map[string]Callback{}
	callbacks["add"] = func() error {
		fmt.Println("add")
		return nil
	}
	callbacks["add"]() //执行函数

}
