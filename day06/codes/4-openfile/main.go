package main

import (
	"fmt"
	"os"
	"time"
)

//文件追加
/*
O_RDONLY 以只读方式打开
O_WRONLY 以只写方式打开
O_RDWR 以读写方式打开
O_APPEND 追加
O_CREATE 文件夹不存在时候创建
O_EXCL  文件必须存在
O_SYNC  使用同步io
O_TRUNC
*/

// os.Open => 读，文件不存在报错
//func Open(name string) (*File, error) {
//	return os.OpenFile(name, os.O_RDONLY, 0777)
//}
//
//// os.Create => 写文件，文件存在 截断 文件不存在 创建
//func Create(name string) (*File, error) {
//	return os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC)
//}

func main() {
	path := "test.log"
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println(err)
		return

	}
	defer file.Close()
	fmt.Fprintf(file, "%s\n", time.Now().Format("2006-01-02 15:04:05"))

}
