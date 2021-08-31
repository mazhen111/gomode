package main

import (
	"fmt"
	"visibility/models"
)

//结构体可见性
func main() {

	publicVar := models.PublicStruct{}
	fmt.Printf("%#v", publicVar)
	//fmt.Println(publicVar.privateAttrPr)\
	publicVar1 := models.NewPrivateStruct()
	publicVar1.PublicAttrPr = "mazhen"
	fmt.Println(publicVar1.PublicAttrPr)
	//简写是小写是可以调用的
	combindVar := models.CombindStruct{}
	fmt.Println(combindVar.PublicStruct.PublicAttrPu)
	fmt.Println(combindVar.PublicAttrPu)
	// fmt.Println(combindVar.privateStruct.PublicAttrPr)
	fmt.Println(combindVar.PublicAttrPr)

}
