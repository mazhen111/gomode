package main

import "fmt"


func main() {
	name := "kk"
	nums := []int{1, 2, 3}
	fmt.Println("111")
	func(pname string,pnums [] int){
		fmt.Println(pname,pnums) // kk, [1, 2, 3]
		pname = "silence"
		pnums = []int{1, 2}
		fmt.Println(pname,pnums) //silence [1 2]
	}(name,nums)
	fmt.Println(name,nums) //kk [1 2 3]
}
