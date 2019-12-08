package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	resp, err := http.Get("https://www.liwenzhou.com/")
	if err != nil {
		fmt.Println("请求出错了,err:",err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read response error:",err)
		return
	}
	fmt.Printf("%T",body)
}
