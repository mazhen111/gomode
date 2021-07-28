package main

import "fmt"

func main() {
	fmt.Println("1111")
	muns := make([]int, 5, 10)
	fmt.Println(len(muns))
	fmt.Println(cap(muns))
	//nums2 := muns[2:5]
	//fmt.Println(nums2 ,len(nums2),cap(nums2))
	nums2 := muns[2:5:5]
	fmt.Println(nums2, len(nums2), cap(nums2))
	nums2 = append(nums2, 1000)
	muns = append(muns, 1)
	fmt.Println(nums2, muns)

}
