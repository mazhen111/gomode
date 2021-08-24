package main

import "fmt"
import (
	 "8-testmultpkg/pkga"
     "8-testmultpkg/pkgb"
)

func main() {
	fmt.Println("11")
	pkga.Test()
	pkgb.Test()

}



