package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

const bufferSize = 1024

func main() {
	var (
		src string
		dst string
	)
	flag.StringVar(&src, "src", "", "src file path")
	flag.StringVar(&dst, "dst", "", "")

	flag.Usage = func() {
		fmt.Println("usage: cp --src path --dst path")
	}
	flag.Parse()
	srcfile, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer srcfile.Close()
	dstfile, err := os.Create(dst)
	if err != nil {
		fmt.Println(err)
		return
	}
	scx := make([]byte, bufferSize)

	for {
		n, err := srcfile.Read(scx)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		dstfile.Write(scx[:n])

	}

}
