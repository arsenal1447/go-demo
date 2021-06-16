package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.zxx123.cn/")
	// resp, err := http.Get("https://www.liwenzhou.com/")
	if err != nil {
		fmt.Println("get failed , err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read from resp.Body failed , err:%v\n", err)
		return
	}
	fmt.Print(string(body))

}
