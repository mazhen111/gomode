package main

import "fmt"

func factV1(a int) int {
	if a < 0{
		return -1
	}else if a == 0 {
		return 1
	}else {
		b:=1
		for i:=1;i<a;i++{
			b *=i


		}
		return b

	}
}

func fact (a int) int {
	if a < 0{
		return -1
	}else if a ==0 {
		return 1
	}else {

		return a *fact(a-1)

	}
	
}

func main() {
	fmt.Println(factV1(10))
	fmt.Println(fact(3))
}
