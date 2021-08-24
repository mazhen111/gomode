package main

//映射(map)

/*
映射是存储一系列无序的 key/value 对，通过 key 来对 value 进行操作（增、删、改、 查）。
映射的 key 只能为可使用==运算符的值类型（字符串、数字、 布尔、数组）， value 可以为
任意类型
 */
import "fmt"

func main() {
	// G -> 名字
	// key type
	// value type
	// map[keytype]valuetype
	var names map[string]string=map[string]string{"test":"ha"}
	fmt.Printf("%T %v\n",names,names)
	//添加元素
	names["test1"]="2222"
	fmt.Println(names)
	//访问元素
	fmt.Println(names["test1"])
	//判断映射为空，还是赋值为空
	names["test2"]=""
	v,ok:=names["test2"] // v = k => value k不存在返回value type 0值， ok bool
	fmt.Println(v,ok)
	v,ok =names["test4"] // v = k => value k不存在返回value type 0值， ok bool
	fmt.Println(v,ok)
	//删除元素
	delete(names,"test2")
	fmt.Println(names)
	//映射遍历
	for i,v :=range names{
		fmt.Println(i,v)
	}
	//make 定义映射
	var scores =make(map[string]int)
	scores["Go3028"] = 80
	fmt.Println(scores)
	scores["Go3027"] = 82
	scores["Go3026"] = 82
	scores["Go3025"] = 82
	scores["Go3029"] = 82
	for i,v :=range scores{
		fmt.Println(i,v)
	}








}
