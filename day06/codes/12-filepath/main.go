package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

//学习filepath包，兼容各操作系统的文件路径
func main() {
	fmt.Println("11")
	//返回所给路径的绝对路径
	path, _ := filepath.Abs("1.txt")
	fmt.Println(path)
	//返回路径最后一个元素
	fmt.Println(filepath.Base(path))
	//如果路径为空字符串，返回.
	fmt.Println(filepath.Base(""))
	//如果路径只有斜线，返回/
	fmt.Println(filepath.Base("///"))
	//分割路径中的目录与文件
	fmt.Println(filepath.Split(path))
	//将路径使用路径列表分隔符分开，见os.PathListSeparator
	fmt.Println(filepath.SplitList("a;b;c"))
	//返回路径中的扩展名
	//如果没有点，返回空
	fmt.Println(filepath.Ext(path))
	fmt.Println(filepath.HasPrefix(path, "d:\\a"))
	//判断路径是不是绝对路径
	fmt.Println(filepath.IsAbs(path), filepath.IsAbs("1.go"))
	//返回所有匹配的文件
	match, _ := filepath.Glob("./*.go")
	fmt.Println(match)
	filepath.Walk("./", func(path string, info fs.FileInfo, err error) error {
		fmt.Println(path, info.Name())
		return nil
	})
	//连接路径，返回已经clean过的路径
	fmt.Println(filepath.Join("C:/a", "/b", "/c"))

}
