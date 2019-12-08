package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main(){
	apiUrl := "http://127.0.0.1:8888/getData?id=1&love=1"
	u, err := url.ParseRequestURI(apiUrl)
	fmt.Printf("%T\n",u)
	if err != nil {
		fmt.Printf("parse url requestUrl failed,err:%v\n", err)
	}
	//解析URI字符串为URL对象，格式不对会报错。
	fmt.Println("Host->",u.Host)
	fmt.Println("Path->",u.Path)
	fmt.Println("RawPath->",u.RawPath)
	fmt.Println("RawQuery->",u.RawQuery)
	fmt.Println("ForceQuery->",u.ForceQuery)
	//Host       string    // host or host:port
	//Path       string    // path (relative paths may omit leading slash)
	//RawPath    string    // encoded path hint (see EscapedPath method)
	//ForceQuery bool      // append a query ('?') even if RawQuery is empty
	//RawQuery   string    // encoded query values, without '?'
	data := url.Values{}
	data.Set("name","jack")
	data.Set("age","7")
	data.Set("age","7")
	data.Add("age","71")
	fmt.Printf("%T\n",data)
	fmt.Println(data.Encode())
	u.RawQuery = data.Encode()
	fmt.Println(u.String())
	resp, err := http.Get(u.String())
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
