package main

import (
	"fmt"
	"sync"
)

var input chan int
var input2 chan int
var wg = sync.WaitGroup{}

func job(input chan <- int){
	defer wg.Done()
	for i := 0;i < 10;i++ {
		//time.Sleep(time.Second*10)
		input <- i
	}
	close(input)
}

func worker(input chan <- int,output <- chan int){
	defer wg.Done()
	defer close(input)
	for x := range output {
	   	input <- x*x
	}
}

func main(){
	input = make(chan int,100)
	input2 = make(chan int,100)
	wg.Add(2)
	go job(input)
	go worker(input2,input)

	for x := range input2 {
		fmt.Println(x)
	}
	wg.Wait()
}
