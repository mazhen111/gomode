
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

	//添加元素
	//切片的末尾添加
	names = append(names,"xiaolong")

	//删除元素
	//切片操作
	// names[start:end] names中从start开始到end-1所有元素组成的切片
	// names[1:10]
	//			0	1								9	 10		11
	// names = ["" "赵昌建" "" "" "" "" "" "" "" "卫智鹏" "sen" "刘冉"]
	// [names[1], names[2], ... names[9]]
	//删除索引为0
	//fmt.Println(names[1:len(names)])
	//删除最后一个元素
	//fmt.Println(names[0:len(names)-1])
	//删除中间的值
	//slice := []int{1, 2, 3, 4, 5, 6, 7}
	//slice =append(slice[:3],slice[5:]...)
	//fmt.Println(slice)
	nums:=[]int{0, 1, 2, 3, 4, 5}
	nums2 := []int{10, 11, 12, 13, 14, 15, 16}
    copy(nums ,nums2)
	fmt.Println(nums,nums2)
	// 切片操作和原来的切片共享存储空间

















}
