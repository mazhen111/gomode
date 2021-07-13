
//切片
package main

import (
	"fmt"
)

func main() {
	//定义names 元素类型为string的切片
	//var names []string //null
	//fmt.Println(names)
	//fmt.Printf("%T\n",names)
	var names=[]string{"ma","zhen"}
	fmt.Println(names)
	fmt.Printf("%T",names)
	//字面量
	//[]type{} => 空切片
    //[]type{v1, v2, ..., vn}
    //// []type{i1:v1, i2:v2, in:vn}
    names = []string{1:"mazhen",3:"ma",7:"hahahha"}
    fmt.Printf("%T\n",names)
    fmt.Printf("%q\n",names)
    //访问 修改元素
    //索引
    fmt.Println(names[1])
    fmt.Println(names[7]) //超出索引，编译不会报错，支持的时候会出错
    //修改
    names[4]="shell"
    fmt.Println(names[4])
    //长度
    fmt.Println(len(names))
    //切片遍历
    for i:=0 ; i<len(names) ;i++{
    	fmt.Println(i,names[i])
	}
	for a,v := range names{
		fmt.Println(a,v)
	}





}
