package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	path := "test.txt"
	file, error := os.Open(path)
	fmt.Println(file, error)
	if error != nil {
		return

	}
	defer file.Close()
	content := make([]byte, 3)
	for {
		n, error := file.Read(content) //n读取的字节
		if error != nil {
			if error != io.EOF {
				fmt.Println(error)

			}
			break
		}

		fmt.Println(string(content[:n]))

	}

}
