package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func readFileWithBufio(path string) io.Reader {
	fr, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	br := bufio.NewReader(fr)
	return br
}

func main() {
	url := "http://127.0.0.1:3000/postData?id=1"

	// 表单数据
	contentType := "image/jpeg"
	//data := "name=小王子&age=18"

	//contentType := "application/json"

	//data := `{"name":"小王子",age:18}`

	br := readFileWithBufio("121.jpg")

	resp, err := http.Post(url,contentType,br)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(b))
}
