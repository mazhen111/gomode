package main
//匿名函数
import (
	"fmt"
	"strings"
)

func print(formatter func(string)string,args...string)  {
	for i,v :=range args{

		fmt.Println(i,formatter(v))

	}

}


func main() {
	fmt.Println("111")
	names:=[]string{"赵昌建", "kk", "17-赵"}
	start := func(txt string) string {
		return "*"+txt+"+"

	}
	print(start,names...)
	//匿名函数
	c:= func() {
		fmt.Println("我是匿名函数")
	}
	c()
	fmt.Println(strings.FieldsFunc("AbaabcAbcabc", func(ch rune) bool {
		return ch == 'a'
	}))





}
