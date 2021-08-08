package main

import "fmt"

//接收可变参数列表 string
//函数格式化输出(针对可变参数列表中的每个元素) string->string
func print(formatter func(string) string,args...string)  {
	for i,v :=range args{

		fmt.Println(i,start(v))
	}
}

func start(txt string) string {
	return "*" + txt + "*"
}



func main() {

	fmt.Println("111")
	names := []string{"赵昌建", "kk", "17-赵"}
	print(start,names...)

}