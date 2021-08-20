package main

import "fmt"
import "7-pkgvar/testpkg"

func main() {
	fmt.Println("11")
	testpkg.T1Func()
	fmt.Println(testpkg.T1Name)
}
