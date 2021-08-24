package main

import "fmt"

func add(ages...int)  {
	for j :=0;j<=len(ages)-1;j++ {
		for i := 0; i < len(ages)-1; i++ {
			if ages[i] > ages[i+1] {
				ages[i], ages[i+1] = ages[i+1], ages[i]

			}

		}
	}
	fmt.Println(ages)



}

func main() {
	nums:=[]int{7,9,5,8,2,1,4}
	add(nums...)


}
