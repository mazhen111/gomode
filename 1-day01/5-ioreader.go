package main

import "fmt"

func main() {
	var name string
	fmt.Println("请输入姓名")
	fmt.Scan(&name) // 接收控制台输入内容 赋值给变量name
	// &name => 取name指针(地址)
	fmt.Printf("输入的内容%s\n",name)
	fmt.Println("输入年龄")
	var age int
	fmt.Scan(&age)
	fmt.Println(age)
}
