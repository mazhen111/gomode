package main

import "fmt"

func test(stop bool)  {
	fmt.Println("a")
	if stop{
		fmt.Println("b")
		return   //返回值并退出
	}

	fmt.Println("c")

}
func main() {
	test(true)
	fmt.Println("--------------")
	test(false)

}