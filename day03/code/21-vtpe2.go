package main

import "fmt"
func main() {
	name := "kk"
	nums := []int{1, 2, 3}
	func(name string ,pnums [] int){
		fmt.Println(name,nums) // kk [1 2 3]
		name="silence"
		nums[0]=100
		fmt.Println(name,nums) //silence [100 2,3]
	}(name,nums)
	fmt.Println(name,nums) //kk [100 2 3]

}
