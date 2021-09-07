package main

import (
	"fmt"
	"os"
)

/*
标准输出流
*/
/*
fmt模块
Errorf：创建 error 类型
Fprint：将数据输出到输出流中，不添加换行
Fprintf：将数据输出到按照格式输出到输出流中
Fprintln：将数据输出到输出流中，并添加换行
Fscan：从输入流中读取数据
Fscanf：从输入流中按照指定格式读取数据
Fscanln：从输入流中读取数据， 回车作为结束符
Print：将数据输出到标准输出流中，不添加换行
Printf：将数据输出到按照格式输出到标准输出流中
Println：将数据输出到标准输出流中，并添加换行
Scan：从标准输入流中读取数据
Scanf：从标准输入流中按照指定格式读取数据
Scanln：从标准输入流中读取数据，回车作为结束符
Sprint：将数据转化为字符串，不添加换行
Sprintf：将数据按照格式转化为格式为字符串
Sprintln：将数据转化为字符串，并添加换行
Sscan：从字符串中读取数据
Sscanf:从字符串中按照格式读取数据
Sscanln:从字符串中读取数据，回车作为结束符
*/
func main() {
	//标准输入 ===>命令行的文件 os.Stdin
	//标准输入 =====> 命令行文件 os.Stdout
	//标准错误输出=====> os.Stderr
	//content := make([]byte,3)
	fmt.Println("请输入内容")
	//os.Stdin
	//fmt.Println(os.Stdin.Read(content))
	//fmt.Ccan 从标准输入流中读取数据
	//fmt.Printf("q\n",string(content))
	os.Stdout.WriteString("我是Stdout的输出")
	fmt.Print("xxxx")
	//fmt.Print / Println / Printf
	//fmt.Sprint / Sprintln / Sprintf
	//Fprint：将数据输出到输出流中，不添加换行
	fmt.Fprint(os.Stdout, "AAAA")
	//Fprintf：将数据输出到按照格式输出到输出流中
	fmt.Fprintln(os.Stdout, "aaaaa")
	//fmt.Fprintf(os.Stdout, "i am: %s", "aaaaa")
	// fmt.Scan => Scanln, Scanf
	// fmt.Sscan => 从字符串扫描到变量
	// fmt.Fscan => 从文件扫描到变量
	name := ""
	//fmt.Println(fmt.Scan(&name))
	//fmt.Printf("%q\n", name)
	// fmt.Println(fmt.Scanln(&name))
	// fmt.Printf("%q\n", name)
	fmt.Println(fmt.Scanln(&name))
	fmt.Println(fmt.Scanf("name%s\n", &name))
	fmt.Printf("%q\n", name)

}
