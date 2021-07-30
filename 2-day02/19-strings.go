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
	fmt.Println(strings.HasPrefix("abc","a"))
	//HasSuffix 判断字符串以什么结尾
	fmt.Println(strings.HasSuffix("AAAAb","b"))
	//将字符串数组（或slice）连接起来可以通过 Join 实现。
	fmt.Println(strings.Join([]string{"1", "2","4"}, "-"))
	//打印20次*
	fmt.Println(strings.Repeat("*", 20))
	// 用 new 替换 s 中的 old，一共替换 n 个。
	// 如果 n < 0，则不限制替换次数，即全部替换
	fmt.Println(strings.Replace("abcabcabcdabc","abc","x",2))
	//全部替换
	fmt.Println(strings.ReplaceAll("abcabcabcdabc","abc","x"))
   //用中划线分割
	fmt.Printf("%q\n", strings.Split("a-b--c", "-"))
	fmt.Printf("%q\n", strings.SplitAfter("a-b-c", "-"))
	fmt.Printf("%q\n", strings.SplitN("a-b-c", "-", 2))
	//第一个字母修改为大写
	fmt.Println(strings.Title("a b bab abc"))
	//ToLower将s中的所有字符修改为其小写格式
	fmt.Println(strings.ToLower("SSSS"))
	//ToUpper将s中的所有字符修改为其大写格式
	fmt.Println(strings.ToUpper("sss"))
	//Trim将删除s首尾连续的包含在cutset中的字符：abc
	fmt.Println(strings.Trim("abcbcdabcdbcabc","abc"))
	//TrimLeft将删除s头部连续的包含在cutset中的字符：abc
	fmt.Println(strings.TrimLeft("abcbcdabcdbcabc","abc"))
	//TrimRight将删除s尾部连续的包含在cutset中的字符：abc
	fmt.Println(strings.TrimRight("abcbcdabcdbcabc","abc"))
	//TrimSpace将删除s首尾连续的的空白字符
	fmt.Println(strings.TrimSpace("abcbcdabcdbcabc \\n\\r   "))
	//TrimSuffix删除s尾部的suffix字符串
	fmt.Println(strings.TrimSuffix("abcbcdabcdbcabc", "abc"))
	//TrimPrefix删除s头部的prefix字符串
	fmt.Println(strings.TrimPrefix("abcbcdabcdbcabc", "abc"))











}

