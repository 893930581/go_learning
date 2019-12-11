package main

import (
	"fmt"
	"html"
	"net/http"
)

func sayHello(w http.ResponseWriter,r *http.Request) {
	defer r.Body.Close()
	fmt.Println("Method",r.Method)
	fmt.Println("URL",r.URL)
	fmt.Println("URL.Host",r.URL.Host)
	fmt.Println("URL.RawQuery",r.URL.RawQuery)
	fmt.Println("Proto",r.Proto)
	fmt.Println("Header",r.Header)
	fmt.Println("ContentLength",r.ContentLength)
	fmt.Println("Host",r.Host)
	fmt.Println("Form",r.Form)
	r.ParseForm()
	fmt.Println("PostForm",r.PostForm)
	fmt.Println("RequestURI",r.RequestURI)
	fmt.Println("MultipartForm",r.MultipartForm)
	w.Header().Add("Set-Cookie","name=1")
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
