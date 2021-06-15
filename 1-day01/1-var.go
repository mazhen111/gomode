package main

import "fmt"
// 我是行注释
/*
我是块注释
 */
//全局变量

/*
var 定义变量
var 标识符（变量名）类型（变量的类型）
string 表示字符串
// 若未设置值则使用零值进行初始化

1. var flag type
2. var flag type = value
	3. var flag = value
	4. flag := value
 */
//全局变量
var name = "kk"
func main() {
	//局部变量
	//var age int = 1+2
	//var weight =30 //申明变量忽略类型，通过直来推导
	height, weight := 168, 138
	/*
	var age, weight, height int
	var age, weight, height int = 1, 2, 3
	var age, weight, height = 1, "", 3
	 */
	fmt.Println(name,height,weight)

}