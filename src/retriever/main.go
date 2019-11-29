package main

import (
	"fmt"
	"go_learning/src/retriever/mock"
	"go_learning/src/retriever/real"
	"time"
)

const url string = "http://www.imooc.com"

type Getter interface {
	Get(url string) string
}

func download(r Getter) string {
	return r.Get(url)
}

type Poster interface {
	Post(url string,form map[string]string) string
}

func Post(poster Poster) {
	poster.Post(url,
		map[string]string {
			"name": "ccmouse",
			"course": "golang",
		})
}

type GetPoster interface {
	Getter
	Poster
}

func session(s GetPoster){
	s.Get(url)
	s.Post(url,map[string]string{
		"contents": "another faked imooc.com",
	})
}


func main() {
	var r Getter
	r = &mock.Http{ Contents : "this is a facebook" }
	fmt.Printf("类型：%T 值：%v\n",r,r)
	inspect(r)
	r = &real.Http{
		//指针接收者方法只能传递指针
		UserAgent: "Mozilla/5.0",
		TimeOut: time.Minute,
	}
	fmt.Printf("类型：%T 值：%v\n",r,r)
	inspect(r)

	//Type Asserion -> what means a Type Asserion?
	realHttp := r.(*real.Http)
	fmt.Println(realHttp.TimeOut)
	if mockHttp,ok := r.(*mock.Http);ok {
		fmt.Println(mockHttp.Contents)
	}

	//fmt.Println(download(r)
}

//Type Switch
func inspect(r Getter) {
	fmt.Printf("类型：%T 值：%v\n",r,r)
	switch v := r.(type) {
	case *mock.Http:
		fmt.Println("Contents:",v.Contents)
	case *real.Http:
		fmt.Println("Contents:",v.UserAgent)
	}
}