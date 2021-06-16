package main

import "fmt"

func main()  {
	sum :=0
	for idex :=0 ; idex <=100; idex ++{
		sum += idex
		fmt.Println(sum)

	}
	fmt.Println(sum)

}