package main

import "os"

var workserQueue = make(chan int, 4) //管道实现限制例程数量
func checkErr(err error) {
	if err != nil {
		panic(err)
	}

}
func isExists(p string) (bool, error) {
	_, err := os.Stat()

}
