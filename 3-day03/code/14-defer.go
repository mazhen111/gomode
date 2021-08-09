package main
import "fmt"
//函数延期执行
func main() {
	fmt.Println("start")
	//defer 函数调用
	//defer 延迟执行
	//在函数退出之前执行
	defer func() {
		fmt.Println("defer")

	}() //返回值

	fmt.Println("end")



}
