package main

import (
	"fmt"
	"html"
	"net/http"
)

func sayHello(w http.ResponseWriter,r *http.Request) {
	fmt.Println("Method",r.Method)
	fmt.Println("URL",r.URL)
	fmt.Println("Proto",r.Proto)
	fmt.Println("Header",r.Header)
	fmt.Println("ContentLength",r.ContentLength)
	fmt.Println("Host",r.Host)
	fmt.Println("Form",r.Form)
	fmt.Println("PostForm",r.PostForm)
	fmt.Println("RequestURI",r.RequestURI)
	fmt.Println("TransferEncoding",r.TransferEncoding)
	defer fmt.Fprintf(w,"hello server run by GO %q!!",html.EscapeString(r.URL.Path))
}

func main () {
	http.HandleFunc("/",sayHello)
	err := http.ListenAndServe(":3000",nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
