package main

import (
	"fmt"
	"math/rand"
	"time"
)

func task(result chan<- int64) {
	interval := rand.Intn(10)
	fmt.Println("sleep", interval)
	time.Sleep(time.Duration(interval) * time.Second) //休眠时间
	result <- time.Now().Unix()

}

func main() {
	rand.Seed(time.Now().Unix())
	var result chan int64 = make(chan int64)
	var timeout chan int = make(chan int)
	go func() {
		time.Sleep(5 * time.Second)
		close(timeout)

	}()
	fmt.Println(time.Now())
	go task(result)
	select {
	case r := <-result:
		fmt.Println("success:", r)
	case <-timeout:
		fmt.Println("timeout")

	}
	fmt.Println(time.Now())
}
