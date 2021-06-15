package main

import "fmt"

func main (){
	fmt.Println("老公的想法")
	// 判断 是否有卖西瓜的 (控制台输入y)
	// 有 买1个包子
	// 否则买10个包子
	var text string
	fmt.Println("有西瓜的买")
	fmt.Scan(&text)
	if text == "y"{
		fmt.Println("1买包子")
	} else {
		fmt.Println("10个包子")


	}


}
