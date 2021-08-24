package main

import (
	"errors"
	"fmt"
	"strconv"
)

//error <===>接口

func dev(lefs,right int) (int ,error) {
	if right ==0{
		return 0 , fmt.Errorf("right num is zero") //初始化error类型变量
	}
	return lefs /right,nil

}
func  divV2(left, right int) (int, error) {
if right == 0 {
return 0, errors.New("right num is zero") //初始化error类型变量
}
return left / right, nil
}


func main() {
	//需要给调试者返回错误信息
	//通过函数最后一个返回值返回错误信息
	//调用者需要对错误进行检查, 决定自己如何操作
	//建议程序员去处理错误
	v,err :=strconv.Atoi("123abcd")
	if err == nil{
		fmt.Println("value是", v)
	}else {
		fmt.Println(err)
	}
	fmt.Println(dev(2,1))
	fmt.Println(divV2(2,1))

}
