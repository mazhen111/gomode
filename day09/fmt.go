package main

import (
	"fmt"
	"strconv"
)

func main() {
	//发送
	prefix := fmt.Sprintf("%05d", 1)
	ctx := []byte(prefix)
	lengthText := string(ctx)
	length, err := strconv.Atoi(lengthText)
	fmt.Println(length, err)

}
