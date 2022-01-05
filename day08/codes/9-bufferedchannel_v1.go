package main

import "fmt"

//带缓冲区管道

func main() {
	var channel chan int = make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			channel <- i

		}
		//close(channel)

	}()

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)

	}

}
