package main

import "fmt"

func test() (i int) {
// 在延迟执行中尽量不要修改返回值
// var i int = 1
i = 1
defer func() {
fmt.Println("defer")
}()

return i
}


func main() {
	fmt.Println(test())

}