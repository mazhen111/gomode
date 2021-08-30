package main

import "fmt"
type Counter =int

func main() {
	fmt.Println("11")
	var counter Counter
	fmt.Printf("%T, %v\n", counter, counter)
	var num int = 10
	fmt.Println(counter + num)

	var (
		r rune
		b byte
	)
	fmt.Printf("%T, %T", r, b)

}

