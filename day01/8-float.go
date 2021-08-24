package main

import "fmt"

func main() {
	height :=1.69
	/*
	%f 、%F ：十进制表示法 ：十进制表示法 ：十进制表示法 ：十进制表示法
	%n.mf 表示最小 表示最小 占 n个宽度并且保留 个宽度并且保留 个宽度并且保留 个宽度并且保留 m位小数 位小数
	%e 、%E ：科学记数法表示 ：科学记数法表示 ：科学记数法表示 ：科学记数法表示
	%g 、%G ：自动选择 ：自动选择 ：自动选择 最紧凑的表示 紧凑的表示 方法 %e(% E) 或%f(%F) %f(%F)%f(%F)
	*/
	fmt.Printf("%T %f %g %e\n",height,height,height,height)
	//运算
	//算数运算
	// + - * / ++ --
	fmt.Println(1.2 + 1.1) //2.3
	fmt.Println(1.2 - 1.1) //0.1
	fmt.Println(1.2 * 1.1) //1.32
	fmt.Println(1.2 / 1.1 ) //1.0909090909090908
	height ++
	fmt.Println(height)
	height --
	fmt.Println(height)
	//关系运算
	// > >= <= <
	fmt.Println(1.1 > 1.2) //false
	fmt.Println(1.1 >= 1.2) //false
	fmt.Println(1.1 < 1.2 ) //true
	fmt.Println(1.1 <= 1.2) //true
	// 赋值预算
	// =, += , -=, /=, *=
	// a += b a = a + b
	// 类型转换 float64(), float32()
	fmt.Printf("%f %10.3f\n", height, height)
}