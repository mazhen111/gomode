package main

import (
	"fmt"
	"io"
	"os"
)

const bufferSiz = 1024

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: tailf path")
		return
	}
	path := os.Args[1]
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	defer file.Close()
	content := make([]byte, bufferSiz)
	for {
		n, err := file.Read(content)
		if err != nil {
			if err == io.EOF {

			} else {

				fmt.Println(err)
				break
			}

		} else {

			fmt.Println(content[:n])
		}

	}

}
