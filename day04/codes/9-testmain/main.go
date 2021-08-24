package main

import "fmt"
import _ "9-testmain/pkg"

/*

main
程序的入口
开发者写代码的执行入口
 */

func main() {

	fmt.Println("22")


}

// 初始化函数
// 导入包时执行

func init() {
	fmt.Println("main.init")

}