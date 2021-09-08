package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("1.txt")
	if err != nil {
		fmt.Println(err)
		return

	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString("1111")
	//把内存内容刷新到文件里
	writer.Flush()

}
