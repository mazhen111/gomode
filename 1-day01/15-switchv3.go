package main

import "fmt"

func main() {
	// 控制台输入一个成绩
	// >= 90 =>A
	// >= 80 => B
	// >= 70 => C
	// >= 60 => D
	// 其他 => E
	var score int
	fmt.Println("请输入考生成绩")
	fmt.Scan(&score)
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 70:
		fmt.Println("C")
	default:
		fmt.Println("D")
	}
}