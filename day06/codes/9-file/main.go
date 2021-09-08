package main

import (
	"fmt"
	"os"
)

func main() {
	path := "test"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	// open file dir 都可以打开
	// file => 读取文件内容
	// dir => 读取目录下的文件名
	// 如何判断是文件 还是 目录
	fileinfo, error := file.Stat()
	fmt.Println(fileinfo, error)
	//获取文件夹名称
	fmt.Println(fileinfo.Name())
	//获取文件权限
	fmt.Println(fileinfo.Mode())
	//获取是否是文件夹
	fmt.Println(fileinfo.IsDir())
	//获取文件夹创建时间
	fmt.Println(fileinfo.ModTime())
	//获取文件夹大小
	fmt.Println(fileinfo.Size())
	//读取文件夹下的内容
	//fmt.Println(file.Readdirnames(-1))
	fileInfos, err := file.ReadDir(-1)
	for _, fileInfo := range fileInfos {
		fmt.Println(fileInfo.Name(), fileInfo.IsDir())
	}
}
