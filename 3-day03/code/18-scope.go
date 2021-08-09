package main
//作用域
import (
	"fmt"
)

func main() {
	name := "kk"
	nums := []int{}
	fmt.Println("11")
	func(){
		fmt.Println(name,nums) // kk, []
		name = "silence"
		nums = []int{1, 2, 3}
		fmt.Println(name,nums) // silence, [1, 2, 3]
	}()
	fmt.Println(name,nums) // silence, [1, 2, 3]
}


