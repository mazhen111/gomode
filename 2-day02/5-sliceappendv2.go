package main

import "fmt"
//当数组超过分配的空间大小时候会重新申请新的内存，两个数组不在共享存储空间
func main() {
	nums :=[]int{1,2,3,4,5}
	nums2 :=nums[1:3]
	nums =append(nums,1000)
	nums2 =append(nums2,100)
	fmt.Println(nums)
	fmt.Println(nums2)

}
