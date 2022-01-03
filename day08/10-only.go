package main

import "fmt"

//只读只写管道
//只写
func in(channel chan<- int) {
	channel <- 1
	channel <- 2
	channel <- 3
}
func out(channel <-chan int) {
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}

func main() {
	var channel chan int = make(chan int, 3)
	var rcchannel <-chan int = channel
	var wchannel chan<- int = channel
	in(wchannel)
	out(rcchannel)
	fmt.Println(<-channel)

}
