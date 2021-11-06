package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var path string
	flag.StringVar(&path, "path", "", "")
	flag.Usage = func() {
		fmt.Println("usage: md5util --path path")
	}
	flag.Parse()
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	file.Close()
	hasher := md5.New()
	io.Copy(hasher, file)
	fmt.Printf("io.copy: %x\n", hasher.Sum(nil))

}
