package main

import "os"

func main() {
	//文件
	//文件创建
	//os.Create("1.txt")
	//// 读取 => os.Open
	////获取属性
	//// 获取属性 os.Open().Stat / os.Stat
	//// 修改属性 => 权限，所属人
	//// os.Chmod()
	//// os.Chown()
	//// 重命名
	//os.Rename("1.txt","b.txt")
	////删除文件
	//os.Remove("b.txt")
	//文件夹操作
	//os.Mkdir("mazhen",os.ModePerm)
	//创建多级目录
	//os.MkdirAll("mazhen/ma/zhen/11",os.ModePerm)
	// 读取 => os.Open
	// 获取属性 os.Open().Stat/ os.Stat
	// 修改属性  => 权限，所属人
	// os.Chmod()
	// os.Chown()
	//重命名
	//os.Rename("mazhen","ma111")
	//删除
	os.RemoveAll("mazhen")

}
