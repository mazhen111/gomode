//数组
package main

import (
	"fmt"
)

/*
数组是具有相同数据类型的数据项组成的一组长度固定的序列，数据项叫做数组的元素，数
组的长度必须是非负整数的常量，长度也是类型的一部分

数组声明需要指定组成元素的类型以及存储元素的数量（长度）。在数组声明后，其长度不
可修改，数组的每个元素会根据对应类型的零值对进行初始化
 */
func main()  {
	//var names [10] int
	//var scores [10] string
	//fmt.Printf("%T\n",names)
	//fmt.Printf("%T\n",scores)
	//fmt.Println(names)
	//fmt.Printf("%q",scores)
    //指定数组长度: [length]type{v1, v2, …, vlength}
	var scores[4] int =[4]int{1,2,3}
	fmt.Printf("%T\n",scores)
	fmt.Println(scores)
	//使用初始化元素数量推到数组长度: […]type{v1, v2, …, vlength}
	bounds :=[...]int{1,2,3,4}
	fmt.Println(bounds)
	//对指定位置元素进行初始化: [length]type{im:vm, …, sin:in}
	nums :=[10]int{1:10,3:30,5:50,8:80} //前面代表下表
	fmt.Println(nums)
	//操作
	//关系运算
	fmt.Println(scores == bounds) //数组的为数必须相同
	//获取数组的长度
	fmt.Println(len(nums))
	//访问值 修改值
	fmt.Println(nums[1])
	//修改数组
	nums[2]=100
	fmt.Println(nums)
	//遍历
	for i:=0 ;i<len(nums);i++{
		fmt.Println(nums[i])
	}
	for i, v := range nums{
		fmt.Println(i,v)
	}
}
