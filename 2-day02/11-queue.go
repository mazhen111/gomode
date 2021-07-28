package main

import "fmt"

/*队列
先进先出
*/
func main() {
	nums := []int{}
	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	nums = append(nums, 4)
	nums = append(nums, 5)
	for len(nums) != 0 {
		fmt.Println(nums[0])
		nums = nums[1:]

	}

}
