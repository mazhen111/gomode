package main

/*
文件写入
*/
import (
	"fmt"
	"os"
)

func main() {
	path := "test.txt"
	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Println(file.Write([]byte("11111ma")))
	fmt.Println(file.Write([]byte("11111ma")))

}
