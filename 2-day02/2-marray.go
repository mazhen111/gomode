package main

import (
	"fmt"
)

func main() {
	//[length]type
	// type [2]int
	// 二维数组

	var ms [3][2] int
	fmt.Printf("%T\n",ms)
	fmt.Println(ms)
	ms = [...][2]int {
	 [2]int{1, 2},
	 [2]int{3, 4},
	 [2]int{5, 6},
	}
	ms[0][1]=30
	fmt.Println(ms)
	for i ,line := range ms{
		fmt.Println(i,line)
		for a,v :=range  line{
			fmt.Println(a,v)
		}
	}
}