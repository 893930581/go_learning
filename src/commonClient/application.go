package main

import (
	"fmt"
	"go_learning/src/commonClient/handler"
	"log"
	"net/http"
	"os"
)

type userError interface {
	error
	Message() string
}

type appHandler func(res http.ResponseWriter,req *http.Request) error

func errWrapper(handler appHandler) func(res http.ResponseWriter,req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		defer func() {
			if r := recover();r != nil {
				log.Printf("Panic Server Error:%v",r)
				http.Error(res,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		err := handler(res,req)
		if err != nil {
			fmt.Println(fmt.Sprintf("Error Handling Request：\n\r%S"),err.Error())
			if userError, ok := err.(userError); ok {
				http.Error(res,
					userError.Message(),
					http.StatusBadRequest)
				return
			}
			code := http.StatusOK
			switch  {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(res,
				http.StatusText(code),
				code)
		}
	}
}


func main(){
	http.HandleFunc("/static/",errWrapper(handler.StaticFileHandler))

	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		panic(err)
	}
}
