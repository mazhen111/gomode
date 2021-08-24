package testpkg

import "fmt"
/*
包内成员可见性
首字母大写包外可见
首字母小写只能胞内可见
 */
var T1Name = "T1"
func T1Func()  {
	fmt.Println("T1 func")
	t3()
}
func t3()  {
	fmt.Println("T3 func")


}

