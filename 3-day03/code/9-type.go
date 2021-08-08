package main

import(
	"fmt"
	"math/rand"
	"strings"
	"time"
)
//返回函数

func aFilelds(split rune) bool {
	if split == 'a'{
		return true
	}
	return false
}

func sayHi(){
	fmt.Println("hi")
}
func sayHello(){
	fmt.Println("sayHello")
	
}

func genFunc() func() {
	
	if rand.Int()%2 ==0{
		return sayHello
		
	}
	return sayHi
	
}

func main() {

	fmt.Println("11")
	fmt.Printf("%q\n", strings.FieldsFunc("abcdefabcabc", aFilelds))
	rand.Seed(time.Now().Unix())
	f:=genFunc()
	f()

}
