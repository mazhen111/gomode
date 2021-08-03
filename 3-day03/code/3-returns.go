package main

import "fmt"
//没有返回值的函数
func sayHello()  {
	fmt.Println("hello")

}

//返回多个值
// a + b, a - b, a * b, a / b
func op(a,b int) (int ,int,int,int) {
	return a + b, a - b, a * b, a / b



}
// 命名返回值
//复杂函数不要用
func optv2(a,b int)(sum, sub, mul, div int){
	sum = a + b
	sub = a - b
	mul = a * b
	div = a / b
	return



}


func main() {
	sayHello()
	a,b,c,d :=op(10,20)
	fmt.Println(a,b,c,d)
	fmt.Println(optv2(3,2))


}