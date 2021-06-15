package main

import "fmt"

func main() {
	fmt.Println("1-我叫kk")//打印环境会自动换行
	fmt.Println("2-我叫kk")
	fmt.Print("3-我叫kk")//打印不会换行
	fmt.Print("3-我叫kk\n")
	var name = "mazhen"
	fmt.Printf("5.我叫%s\n",name)//通过占位定义标量填充
}
