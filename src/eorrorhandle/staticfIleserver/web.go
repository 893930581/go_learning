package main

import (
	"go_learning/src/eorrorhandle/staticfIleserver/filelisten"
	"net/http"
)

func main() {
	http.HandleFunc("/list/",filelisten.HandleFileList)
	err := http.ListenAndServe(":8888",nil)
	if err != nil {
		panic (err)
	}
}
