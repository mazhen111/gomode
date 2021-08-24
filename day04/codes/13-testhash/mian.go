package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func main() {
	txt := "i am kk"
	md5value := fmt.Sprintf("%X",md5.Sum([]byte(txt)))
	sha1value :=fmt.Sprintf("%X",sha1.Sum([]byte(txt)))
	sha256vaue :=fmt.Sprintf("%X",sha256.Sum256([]byte(txt)))
	sha512vaue :=fmt.Sprintf("%X",sha512.Sum512([]byte(txt)))
	fmt.Println(md5value)
	fmt.Println(sha1value)
	fmt.Println(sha256vaue)
	fmt.Println(sha512vaue)
	//大字符串算MD5
	sha1Hasher :=sha1.New()
	sha1Hasher.Write([]byte("i"))
	sha1Hasher.Write([]byte("am"))
	sha1Hasher.Write([]byte("kk"))
	fmt.Printf("%X",sha1Hasher.Sum(nil))







}

