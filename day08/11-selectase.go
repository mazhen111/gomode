package main

import (
	"fmt"
	"time"
)

func main() {
	//在管道中只要有一个操作成功就执行相关逻辑
	channe1A := make(chan int)
	channe1B := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		close(channe1A)
	}()
	select {
	case v, ok := <-channe1A:
		fmt.Println("a", v, ok)
	case v, ok := <-channe1B:
		fmt.Println("b", v, ok)
	default:
		fmt.Println("default")
	}

}
