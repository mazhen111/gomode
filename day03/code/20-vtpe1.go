package main

import "fmt"
func main() {
	name := "kk"
	nums := []int{1, 2, 3}
	func(name string ,pnums [] int){
		fmt.Println(name,nums) // kk [1 2 3]
		name="silence"
		nums = []int{1, 2}
		fmt.Println(name,nums) //silence [1 2]
	}(name,nums)
	fmt.Println(name,nums) //kk [1 2 3]

}
