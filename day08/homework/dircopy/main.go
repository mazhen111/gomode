package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"sync"
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

func fsrc(p string) (f *os.File, err error) {
	f, err = os.Open(p)
	if err != nil {
		return f, err
	}
	return f, nil

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
func docopy(src, des *os.File, wg *sync.WaitGroup) {
	_ = <-workserQueue
	var b = make([]byte, 1024)
	defer func() {
		_ = src.Close()
		_ = des.Close()
	}()
	for {
		n, err := src.Read(b)
		if err == io.EOF {
			wg.Done()
			return

		}
		_, werr := des.Write(b[:n])
		if werr != nil {
			fmt.Println(werr)
			wg.Done()
			return
		}
	}
}
func dircopy(src, dest string, wg *sync.WaitGroup) {
	sinfo, err := os.Stat(src)
	checkErr(err)
	//检查是否目录
	if sinfo.IsDir() {
		err = os.Mkdir(dest, os.ModeDir)
		if err != nil {
			fmt.Println(err)
			return
		}
		f, err := os.Open(src)
		defer func() {
			_ = f.Close()

		}()
		checkErr(err)
		fnames, err := f.Readdirnames(0)
		checkErr(err)
		for _, v := range fnames {
			np := path.Join(src, v)
			dests := path.Join(dest, v)
			dircopy(np, dests, wg)

		}
	} else {
		sf, serr := fsrc(src)
		checkErr(serr)
		df, derr := fdest(dest)
		checkErr(derr)
		workserQueue <- 1
		wg.Done()
		wg.Add(1)
		go docopy(sf, df, wg)
	}
}
func main() {
	var d = flag.String("d", "", "目标路径")
	var s = flag.String("s", "", "源文件路径")
	flag.Parse()
	var wg sync.WaitGroup
	dircopy(*s, *d, &wg)
	wg.Wait()
}
