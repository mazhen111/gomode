package main

//函数类型

import "fmt"

func mul(a,b int) int {

	return a * b

}

func add(a, b int) int {
	return a + b
}

func main() {

fmt.Println("111")
	var f func(int, int) int = add

	var fs []func(int, int) int

	fs = append(fs, add, mul)

	c := f(2, 3)
	fmt.Println(c)
	for _, f := range fs {
		fmt.Println(f(2, 3))
	}


}
