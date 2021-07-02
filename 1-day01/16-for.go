package main

import "fmt"

func main()  {
	sum :=0
	for idx :=1 ; idx <100 ; idx++ {
		sum += idx
		fmt.Println(sum)
	}
	fmt.Println(sum)

}