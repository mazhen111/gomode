package main

import "fmt"

func main()  {

	/*
	1+2+3+4+5.....+100
	 */
	sun :=0
	idx :=1
	for {
		sun +=idx
		idx ++
		if idx >100{
			break

		}
		fmt.Println(sun)

	}


}
