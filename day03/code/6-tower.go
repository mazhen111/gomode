package main

import (
	"fmt"
)
/*
汉诺塔:
		n个盘子 (开始)start -> (终止)end (借助)layer
		n-1个盘子start -> tmp 借助end
		第N个 start -> end
		n-1个盘子 tmp -> end 借助A
        start => 开始
		tmp => 借助
		end => 终止
		layer => 盘子的数量
 */

func towr(start ,tmp ,end string ,layer int )  {
	if layer == 1{
		fmt.Println(start, "->", end)
		return

	}
	towr(start,end,tmp,layer-1)
	fmt.Println(start, "->", end)
	towr(tmp,start,end,layer-1)

}
func main()  {
	fmt.Println("111")
	towr("a","b","c",10)
}

