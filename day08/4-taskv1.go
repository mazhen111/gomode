package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wgt sync.WaitGroup
	wgt.Add(2)

	go func(wgt *sync.WaitGroup) {

		for i := 'A'; i <= 'Z'; i++ {
			fmt.Printf("%c\n", i)
		}
		defer wgt.Done()

	}(&wgt)

	func(wgt *sync.WaitGroup) {
		for i := 1; i <= 10; i++ {
			fmt.Printf("%d\n", i)
			runtime.Gosched()
		}

		defer wgt.Done()
	}(&wgt)
	wgt.Wait()

}
