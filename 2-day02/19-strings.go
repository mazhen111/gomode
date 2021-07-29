package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("111")
	// Contains 判断字符串 s 中是否包含子串 substr
	fmt.Println(strings.Contains("abv","b"))
	//ContainsAny 判断字符串 s 中是否包含 chars 中的任何一个字符
	fmt.Println(strings.ContainsAny("abc","bc"))
	// Count 计算字符串 sep 在 s 中的非重叠个数
	fmt.Println(strings.Count("bcabdadb","db"))
	// Fields 以连续的空白字符为分隔符，将 s 切分成多个子串，结果中不包含空白字符本身
	fmt.Printf("%q\n", strings.Fields("abc abc\t123\n1\r1"))
	// HasPrefix 判断字符串 s 是否以 prefix 开头




}

