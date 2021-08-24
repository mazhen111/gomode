package main

import "fmt"

func main() {
	//指针
	//值类型
	//赋值 原有的数据复制一份给新的变量
	// 两个变量之间没有任何联系
	// 对nums进行任何修改都不会影响nums2
	// 对nums2进行任何修改也不会影响nums
	// 有联系 => 引用类型(内部使用使用了指针/地址)
	mums :=[5]int{1,2,3,4,5}
	mums2 :=mums
	fmt.Println(mums,mums2)
	mums2[0]=100
	fmt.Println(mums,mums2)
	var age =1
	var agePointrt *int =nil
	agePointrt = &age // 取出age的地址赋值给agePointer => 取引用
	fmt.Println(agePointrt, age)
	*agePointrt = 33
	fmt.Println(age, *agePointrt)

	var numsPoint *[5]int = &mums
	fmt.Printf("%T\n", numsPoint)
	numsPoint[0] = 100
	fmt.Println(mums, numsPoint)
	

}