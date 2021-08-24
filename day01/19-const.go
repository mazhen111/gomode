package main

import "fmt"

const (

	statusNew int =1
	statusDeleted     = 2
)

func main()  {
	/*
	iota 生成器用于初始化一系列相同规则的常量， 批量声明常量的第一个常量使用
	iota 进行赋值，此时 iota 被重置为 0，其他常量省略类型和赋值，在每初始化一个常量则
	加 1
	 */
	const (
		Monday = 10
		/* 在一个小括号内
		若为赋值，则使用最近的一个已赋值的常量对应的值进行初始化
		*/
		Tuesday = 20
		Wednesday

	)
	fmt.Println(statusNew, statusDeleted)
	fmt.Println(Monday, Tuesday, Wednesday)
}
