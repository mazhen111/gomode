package main

import "fmt"

//堆栈 先进后出
func main() {
	stack := []int{}
	stack = append(stack, 1)
	stack = append(stack, 2)
	stack = append(stack, 3)
	stack = append(stack, 4)
	stack = append(stack, 5)
	fmt.Println(stack)
	for len(stack) != 0 {
		fmt.Println(stack[len(stack)-1])
		stack = stack[:len(stack)-1]

	}

}
