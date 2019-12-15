package main

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		fmt.Printf("%T",r)
		if err, ok := r.(error); ok {
			fmt.Println("Error occured",err)
		}else{
			panic(r)
		}
	}()
	fmt.Errorf()
	b := 0
	a := 5/b
	fmt.Print(a)
}

func main() {
	tryRecover()
}