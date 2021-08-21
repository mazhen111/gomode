package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	fmt.Println(now) //获取当前时间
	fmt.Println(now.Unix()) //把时间转换成时间戳
	/*
	模板占位
	年 2006
	月 01
	日 02
	时 03 （12h） 15 （24h）
	分 04
	秒 05
	*/
	// 年-月-日 24时:分:秒
	format :="2006-01-02 15:04:05"
	times := now.Format(format)
	fmt.Println(times)
	//2018-01-02 16:05 - 年-月-日 (24h)时:分 -> 2006-01-02 15:04
	t1, err := time.Parse("2006-01-02 15:04:05",times)
	fmt.Println(t1,err)




}


