package main

import "fmt"

func main() {
	fmt.Println("int")
	var age int =31
	var char byte ='a'  // 97 => ascii编码 a => 整数值
	var codepoint rune = '我'  // unicode编码 我 => 整数值
	fmt.Println(age,char,codepoint)
	fmt.Printf("%T %T %T\n",age,char,codepoint)
	/*
	%b ：二进制 ：二进制
    %c ：字符 ：字符
	%d ：十进制 ：十进制
	%+d 表示 对正整数 对正整数 对正整数 带+符号
	%nd 表示 最小 占位 n个宽度且右对齐 个宽度且右对齐 个宽度且右对齐
	%-nd 表示 最小 占位 n个宽度且左对齐 个宽度且左对齐 个宽度且左对齐 个宽度且左对齐
	%0nd 表示 最小 占位 n个宽度且右对齐 个宽度且右对齐 个宽度且右对齐 个宽度且右对齐 ，空字符使用 空字符使用 空字符使用 空字符使用 0填充
	%o ：八进制 ：八进制 ，%#o %#o带 0的前缀 的前缀
	%x 、%X ：十六进制 ：十六进制 ：十六进制 ,%#x(%#X) ,%#x(%#X) ,%#x(%#X),%#x(%#X) 带 0x(0X)0x(0X) 0x(0X)的前缀 的前缀
	%U : Unicode : Unicode : Unicode码点 ，%#U 带字符的 带字符的 带字符的 Unicode Unicode 码点
	%q ：带单引号 ：带单引号 ：带单引号 的字符
	*/
	// 十进制 八进制 十六进制
	var (
		t10 = 10
		t8 = 012
		t16 = 0XA
	)
	fmt.Println(t10,t8,t16)
	fmt.Printf("%d %x %o\n",t10,t10 ,t10)
	//算数运算
	// + - * / %
	fmt.Println(1+2) //3
	fmt.Println(1-2) //-1
	fmt.Println(1*2) //2
	fmt.Println(1/2) //0
	fmt.Println(1%2) //1
	// 自增 ++ 自减 --
	age ++
	fmt.Println(age)
	age --
	fmt.Println(age)
	//  关系运算
	// < > <= >= != ==
	fmt.Println(1 > 2 ) //false
	fmt.Println(1 >= 2) //false
	fmt.Println(1 < 2) //true
	fmt.Println(1<=2) //true
	fmt.Println(1 != 2) //true
	fmt.Println( 1 == 2) //false
	// 位运算
	// & | ^ << >> &^
	// 两个整数转为二进制进行计算 对应的位进行计算
	// & 两个为1 => 1
	// | 只要有一个为1 => 1
	// ^ 抑或 相同->0, 不同->1
	// << 10 => 2 ^ n
	// >> 0 => /2^n
	// &^ a &^ b 擦除 b(1->a->0)
	//赋值
	// = += -= *= /= %= &= |= ^= <<= >>=
	// a += b => a = a + b
	// a -= b => a = a - b
	a :=10 //int
	a+=5
	fmt.Println(a)
	a -=8
	fmt.Println(a)
	a *= 20
	fmt.Println(a)
	//不同的数据类型之间不能进行运算
	//强制转换类型
	var b int8 =5
	fmt.Println(a+int(b))








}
