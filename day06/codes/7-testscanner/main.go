package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//带缓冲的io

func ScanInt() (int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return strconv.Atoi(scanner.Text())
	}
	return 0, scanner.Err()
}

func main() {
	//scanner:=bufio.NewScanner(os.Stdin)
	//从标准输入扫描文本，读取由空白符分隔的值保存到传递给本函数的参数中，换行符视为空白符
	//for scanner.Scan(){
	//	fmt.Println(scanner.Text())
	//}

	num, err := ScanInt()
	fmt.Println(num, err)

	num, err = ScanInt()
	fmt.Println(num, err)

	num, err = ScanInt()
	fmt.Println(num, err)

	num, err = ScanInt()
	fmt.Println(num, err)

	num, err = ScanInt()
	fmt.Println(num, err)
}
