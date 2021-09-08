package main

import (
	"fmt"
	"os"
)

func main() {
	//fmt.Println("11")
	path := "1.txt"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

}
