package main

import "fmt"

func main() {
	fmt.Println("11")
	nums := []int{1, 2, 3}
	nums2 := []int{3, 4, 5}
	// 解包 => 切片操作
	nums = append(nums, nums2...)
	fmt.Println(nums)
	nums = append(nums[:2], nums[3:]...)
	fmt.Println(nums)

}
