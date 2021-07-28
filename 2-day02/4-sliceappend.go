package main

import "fmt"


/*
cap()函数返回的是数组切片分配的空间大小

当数组超过分配的空间大小时候会重新申请新的内存，两个数组不在共享存储空间
 */
func main() {

	nums :=[]int{1,2,3,4,5}
	nums2 :=nums[1:3]
	fmt.Println(cap(nums2))
	fmt.Println(nums)
	nums2 = append(nums2,111)
	nums = append(nums,5555)
	fmt.Println(nums)
	fmt.Println(nums2)
}