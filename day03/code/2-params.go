package main
/*
函数用于对代码块的逻辑封装，提供代码复用的最基本方式
 */
import "fmt"

//定义hello
//无参，无返回值
func sayHell() {

	fmt.Println("hello")
}
//有参数，无返回值
func sayHi(name string) {
	fmt.Println("hi",name)

}
func add(a int ,b int) int {

	fmt.Println(a,b)
	return a +b
}


func main() {
	//调用函数，函数名称（参数实参）
	sayHell() //主要小括号
	name:="mazhen"
	sayHi(name)
	c := add(1,2)
	fmt.Println(c)

}

