package main

import (
	"fmt"
	"sync"
)

func main() {
	var wgt sync.WaitGroup
	//使用for循环启动3个例程
	fmt.Println("start") //0,1,2==》3，3，3
	for i := 0; i < 3; i++ {
		wgt.Add(1)
		go func(i int) {
			fmt.Println(i)
			wgt.Done()
		}(i)

	}
	wgt.Wait()

}
