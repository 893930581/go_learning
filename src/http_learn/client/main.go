package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main(){
	client := &http.Client{
	}
	//resp, err := client.Get("http://example.com")
	// ...
	data := "name=小王子&age=18"

	req, _ := http.NewRequest("POST", "http://example.com", strings.NewReader(data))
	// ...

	fmt.Println(req.GetBody())

	req.PostForm.Add("love","you")

	req.Header.Add("If-None-Match", `W/"wyzzy"`)

	_, _ = client.Do(req)

	//resp, err := http.Get("http://127.0.0.1:3000/getData")
	//if err != nil {
	//	panic(err)
	//}
	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	////body, err := ReadFrom(resp.Body,20)
	//fmt.Printf("%T -> %v",body,body)
	//fmt.Println(string(body))
}