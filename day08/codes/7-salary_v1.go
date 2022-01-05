package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var salary int = 0
	var wg sync.WaitGroup
	//定义锁
	var locker sync.Mutex
	fmt.Println("start")
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				locker.Lock() //加锁
				salary += 10
				locker.Unlock()
				runtime.Gosched() //解锁
			}

		}()

		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				locker.Lock() //加锁
				salary -= 10
				locker.Unlock()
				runtime.Gosched() //解锁
			}

		}()

	}
	wg.Wait()
	fmt.Println("end:", salary)

}
