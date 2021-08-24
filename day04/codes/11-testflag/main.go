package main

import (
	"flag"
	"fmt"
)

func main() {
	//ssh -p port [22] root@host
	var port int
	var help bool
	//指定变量与命令参数（通过参数名称）中的关系
	flag.IntVar(&port,"P",22 ,"port")
	flag.BoolVar(&help,"h",false ,"help")
	//解析命令参数
	flag.Parse()
	if help{

		fmt.Println("usage:ssh -p 22 root@locahost")
		flag.PrintDefaults()
		return

	}
	fmt.Println(port)
	//未指定命令参数
	fmt.Println(flag.Args())

}

