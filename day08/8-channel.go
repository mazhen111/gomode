package main

import (
	"fmt"
	"time"
)

func main() {
	//chan
	//chnnel
	//管道中放什么类型的需要提前声明
	//声明int类型的管道channel
	//var channel chan int =make(chan int)
	channel := make(chan int) //声明无缓冲的管道
	//初始化%赋值
	//make
	//channel =make (chan int)
	//操作
	//读与写
	go func() {
		fmt.Println("go start")
		channel <- 1 //将1写入管道
		fmt.Println("go 1 end", time.Now())
		channel <- 2 //将2写入管道
		fmt.Println("go 2 end", time.Now())
	}()
	fmt.Println("channel begin")
	// <-channel // 读管道
	// num := <-channel
	// func test()  int {return 1}
	// test()
	// num := test()
	num, ok := <-channel
	fmt.Println(num, ok)
	time.Sleep(3 * time.Second)
	<-channel
	time.Sleep(2 * time.Second)
	//关闭管道
	//close(channel)
	//go func() {
	//	fmt.Println(<-channel)
	//
	//}()
	////泳道关闭不能写只能读
	////channel <- 1
	//v, ok := <-channel
	//fmt.Println(v, ok)
	go func() {
		channel <- 1
		channel <- 2
		channel <- 3
		close(channel)
	}()
	for num := range channel {
		fmt.Println(num)

	}

}
