package main

import (
	"fmt"
	"go_learning/src/eorrorhandle/staticfIleserver/filelisten"
	"log"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter,request *http.Request) error

func errWraper(handler appHandler) func(writer http.ResponseWriter,request *http.Request){
	return func(writer http.ResponseWriter, request *http.Request) {
 		defer func() {
 			//防止引用程序内部出错。
 			if r := recover(); r != nil {
 				log.Printf("Panic:%v",r)
 				http.Error(writer,
 					http.StatusText(http.StatusInternalServerError),
 					http.StatusInternalServerError)
			}
		}()
		err := handler(writer,request)
		//实际业务执行语句体
		if err != nil {
			fmt.Println(fmt.Sprintf("Errot handling Request：%s"),err.Error())

			if userErr, ok := err.(userError); ok {
				http.Error(writer,
					userErr.Message(),
					http.StatusBadRequest)
				return
			}

			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsNotExist(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(
				writer,
				http.StatusText(code),
				code,)
		}
	}
}

type userError interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/",errWraper(filelisten.HandleFileList))
	err := http.ListenAndServe(":8888",nil)
	if err != nil {
		panic (err)
	}
}
