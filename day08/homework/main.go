package main

import (
	"fmt"
	"os"
)

var workserQueue = make(chan int, 4) //管道实现限制例程数量
func checkErr(err error) {
	if err != nil {
		panic(err)
	}

}
func isExists(p string) (bool, error) {
	_, err := os.Stat(p)
	if err == nil {
		return true, err
	}
	if os.IsExist(err) {
		return false, err

	}
	return false, err
}
func fdest(p string) (f *os.File, err error) {
	if b, _ := isExists(p); b {
		var s string
		fmt.Printf("%s已经存在,是否覆盖[y/n]:", p)
		_, _ = fmt.Scanln(&s)
		if s == "y" || s == "Y" {
			f, err := os.OpenFile(p, os.O_TRUNC|os.O_WRONLY, os.ModePerm)
			return f, err
		}
		if s == "n" || s == "N" {
			return nil, fmt.Errorf("cancel")
		}

	}
	f, err = os.Create(p)
	return f, err
}
