package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {

	/*
	需求：
			猜数字 生成随机数字0-100
			从控制台数据 与生成数字比较
			大 提示太大了
			小 提示太小了
			等于 成功, 程序结束
			最多猜测五次，未猜对，说太笨了，程序结束

	 */
	var randNum int
	rand.Seed(time.Now().Unix())
	randNum = rand.Intn(100)
	for num:=1 ; num<=5;num++{
		fmt.Println(randNum)
		var guessnum int
		fmt.Println("请输入数字")
		fmt.Scan(&guessnum)
		if guessnum > randNum{
			fmt.Println("选大了")

		}else if guessnum < randNum {
			fmt.Println("选小了")
		}else {
			fmt.Println("选择对了")
			break
		}
		if num ==5 {

			fmt.Println("未猜对，说太笨了，程序结束")
		}


	}

}