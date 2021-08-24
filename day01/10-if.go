package main

import "fmt"

func main() {
	fmt.Println("老婆的想法")
	fmt.Println("10个包子")
	var  text string
	fmt.Println("有卖西瓜的吗")
	fmt.Scan(&text)
	if text == "y"{
		fmt.Printf("一个西瓜")
	}
}