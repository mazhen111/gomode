package main

import "fmt"

func test(p bool)  {
	defer func() {
		if err :=recover(); err != nil {
			fmt.Println("出错了, 进行恢复", err)
		}
		
	}()
	if p {
		panic("我是一个错误")
	}

	fmt.Println("running...")
}


func main() {
	fmt.Println("11")
	test(true)

}

