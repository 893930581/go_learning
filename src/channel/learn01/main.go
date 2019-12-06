package main

import (
	"fmt"
	"sync"
)

func g1(i int){
	fmt.Print(i)
}

func deathlock(){
	var a chan int
	a <- 1
	fmt.Print(a)
}

func buffChannel(){
	b := make(chan int,10)
	b <- 10
	b <- 20
	fmt.Println(b)
	c := <- b
	fmt.Println(c)
	close(b)
}

var b chan int

var wg sync.WaitGroup

func main(){
	//buffChannel()
	fmt.Println(b)
	b = make(chan int,10)
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <- b
		fmt.Println(x)
		fmt.Println("在通道中拿到了1")
	}()

	b <- 1
	fmt.Println("1已经被发送到通道中了。")
	wg.Wait()
}
