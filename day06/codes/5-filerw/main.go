package main

import (
	"fmt"
	"os"
)

//光标使用
func main() {
	file, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	file.WriteString("1234")
	file.Seek(-2, 1)
	content := make([]byte, 10)

	fmt.Println(file.Read(content))
	fmt.Println(string(content))
	//移动光标到初始位置
	file.Seek(0, 0)
	file.WriteString("qw")

}
