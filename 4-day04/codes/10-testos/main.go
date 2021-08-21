package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)
func main() {
	num := os.Args
	nums:=0
	fmt.Println(num[1:])
	if len(num) < 3{
		fmt.Println("至少输入两个数字")
		os.Exit(0)

	}
	for i:=0;i <len(num);i++{
		time.Sleep(20)
		if i>=1 {
			 e,err :=strconv.Atoi(num[i])
			 if err == nil {
			 	nums+=e

			 }else {
			 	fmt.Println("输入不是数字格式的字符串时，提示 程序的使用方法为add 1 2 [3 1114 5666]")
			 	break

			 }

			}
			fmt.Println(nums)


		}


}






