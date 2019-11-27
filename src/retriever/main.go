package main

import (
	"fmt"
	"go_learning/src/retriever/mock"
	"go_learning/src/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"this is a facebook"}
	r = real.Retriever{}
	fmt.Println(download(r))
}