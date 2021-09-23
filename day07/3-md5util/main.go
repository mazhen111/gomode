package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"os"
)

const bufferSize = 1024

func main() {
	var path string
	flag.StringVar(&path, "path", "", "")
	flag.Usage = func() {
		fmt.Println("usage: md5util --path path")
	}
	flag.Parse()
	pathfile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	pathfile.Close()
	hasher := md5.New()
	ctx := make([]byte, bufferSize)
	for {
		n, err := pathfile.Read(ctx)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break

		}
		hasher.Write(ctx[:n])

	}
	fmt.Println("%x\n", hasher.Sum(nil))
}
