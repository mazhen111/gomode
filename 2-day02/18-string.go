package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var name string= "kk"
	var desc string = "I love china"
	// 字面量 "" ``
	// 零值 ""
	// 操作
	// 连接 +
	// 关系运算 == != > < >= <=
	// 赋值操作 = +=
	// 长度 len 字节长度
	// 索引 字节
	// 切片 => 字符串 字节
	fmt.Println(name, desc)
	//切片
	fmt.Println(desc[1:4]) //按字节切片
	var txt = "我爱中国"
	//遍历
	for i,v :=range txt{
		fmt.Printf("%d %q \n",i,v)
	}
	//如何计算UTF-8的字符个数
	fmt.Println(utf8.RuneCountInString(txt))
	fmt.Println(utf8.RuneCountInString(desc))
	// 字符串 => []byte
	fmt.Println([]byte(desc))
	//[]byte > 字符串
	fmt.Println(string([]byte(desc)))
	// 字符串 => []rune
	fmt.Println([]rune(desc))
	//[]run > 字符串
	fmt.Println(string([]rune(desc)))









}
