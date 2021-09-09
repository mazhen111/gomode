package main

import (
	"fmt"
	"io"
	"io/ioutil"
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
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	//ReadAll这个函数，用来从io.Reader中一次读取所有数据。
	con, _ := ioutil.ReadAll(file)
	fmt.Println(string(con))
	/*readFile 从filename指定的文件中读取数据并返回文件的内容。成功的调用返回的err为nil而非EOF。
	因为本函数定义为读取整个文件，它不会将读取返回的EOF视为应报告的错误。*/
	fmt.Println(ioutil.ReadFile("1.txt"))
	//ioutil中提供了一个方便的函数：ReadDir，它读取目录并返回排好序的文件和子目录名（[]os.FileInfo）
	fmt.Println(ioutil.ReadDir("A"))
	//WriteFile 将data写入filename文件中，当文件不存在时会根据perm指定的权限进行创建一个,
	//文件存在时会先清空文件内容。对于perm参数，我们一般可以指定为：0666，具体含义os包中讲解。
	ioutil.WriteFile("1.txt", []byte("11111111"), os.ModePerm)
	//复制
	io.Copy(os.Stdout, file)

}
