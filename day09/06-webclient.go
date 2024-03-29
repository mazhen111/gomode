package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	//模拟浏览器
	response, err := http.Get("http://localhost:8888/help")
	//post
	//Head
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response.Proto)       //HTTP/1.1
	fmt.Println(response.StatusCode)  //获取返回值
	io.Copy(os.Stdout, response.Body) //help
	response, err = http.Get("https://www.conardli.top/blog/article/%E5%8D%9A%E5%AE%A2%E6%90%AD%E5%BB%BA/%E3%80%90%E5%8D%9A%E5%AE%A2%E6%90%AD%E5%BB%BA%E3%80%91%E4%B8%AA%E4%BA%BA%E5%8D%9A%E5%AE%A2%E6%90%AD%E5%BB%BA%E5%8F%8A%E9%85%8D%E7%BD%AE.html")
	if err != nil {
		fmt.Println(err)
	} else {
		io.Copy(os.Stdout, response.Body)

	}
	//省略非正常连接不是私密连接
	req, err := http.NewRequest("GET", "https://www.conardli.top/blog/article/%E5%8D%9A%E5%AE%A2%E6%90%AD%E5%BB%BA/%E3%80%90%E5%8D%9A%E5%AE%A2%E6%90%AD%E5%BB%BA%E3%80%91%E4%B8%AA%E4%BA%BA%E5%8D%9A%E5%AE%A2%E6%90%AD%E5%BB%BA%E5%8F%8A%E9%85%8D%E7%BD%AE.html", nil)
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	responses, err := client.Do(req)
	if err != nil {
		fmt.Println(err)

	} else {
		io.Copy(os.Stdout, responses.Body)

	}

}
