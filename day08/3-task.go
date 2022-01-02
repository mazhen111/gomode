package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wgt sync.WaitGroup
	wgt.Add(2)

	go func() {

		for i := 'A'; i <= 'Z'; i++ {
			fmt.Printf("%c\n", i)
		}
		defer wgt.Done()

	}()

	func() {
		for i := 1; i <= 10; i++ {
			fmt.Printf("%d\n", i)
			runtime.Gosched()
		}
	}()
	wgt.Wait()

}
