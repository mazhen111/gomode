package main

import (
	"fmt"
	"runtime"
	"sync"
)

func taskA(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
		//time.Sleep(time.Microsecond * 1000)
		runtime.Gosched()

	}
}
func taskB(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 'A'; i <= 'Z'; i++ {
		fmt.Printf("%c\n", i)
		//time.Sleep(time.Microsecond * 1000)
		runtime.Gosched()
	}

}

func main() {
	fmt.Println("11")
	//time.Sleep(5 * time.Second)
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go taskA(wg) // 启动一个例程
	go taskB(wg)
	wg.Wait()
	fmt.Println("end")

}
