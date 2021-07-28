package main

import "fmt"

func main() {
	//make  初始化数组
	//make make(type, len)
	// 3个参数: make(type, len, cap)
	nuns :=make([]int,3)
	fmt.Println(len(nuns),cap(nuns))
	nums2 := make([]int, 2, 5)
	fmt.Println(len(nums2), cap(nums2))
	fmt.Println(nums2)
	num3:=nums2
	num3=append(num3 ,3)
	fmt.Println(num3)
	nums2 =append(nums2,4)
	fmt.Println(num3)       // [0, 0, 4]
	fmt.Println(nums2)





}